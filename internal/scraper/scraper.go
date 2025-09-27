package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"

	"rugby-promiedos/internal/models"
)

type Scraper struct {
	client *http.Client
}

func NewScraper() *Scraper {
	return &Scraper{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *Scraper) ScrapeESPN(db *gorm.DB) error {
	log.Println("Starting ESPN rugby data scraping...")

	url := "https://www.espn.com/rugby/scoreboard"

	resp, err := s.client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch ESPN rugby page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ESPN returned status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	matches := s.extractMatches(doc)

	for _, match := range matches {
		var existingMatch models.Match
		result := db.Where("url = ?", match.URL).First(&existingMatch)

		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&match).Error; err != nil {
				log.Printf("Failed to save match %s vs %s: %v", match.HomeTeam, match.AwayTeam, err)
			} else {
				log.Printf("Saved new match: %s vs %s", match.HomeTeam, match.AwayTeam)
			}
		} else if result.Error == nil {
			existingMatch.HomeScore = match.HomeScore
			existingMatch.AwayScore = match.AwayScore
			existingMatch.Status = match.Status
			existingMatch.UpdatedAt = time.Now()

			if err := db.Save(&existingMatch).Error; err != nil {
				log.Printf("Failed to update match %s vs %s: %v", match.HomeTeam, match.AwayTeam, err)
			} else {
				log.Printf("Updated match: %s vs %s", match.HomeTeam, match.AwayTeam)
			}
		}
	}

	log.Printf("Scraping completed. Processed %d matches", len(matches))
	return nil
}

func (s *Scraper) extractMatches(doc *goquery.Document) []models.Match {
	var matches []models.Match

	doc.Find("article.scoreboard.rugby").Each(func(i int, sel *goquery.Selection) {
		match := s.parseMatchElement(sel)
		if match.HomeTeam != "" && match.AwayTeam != "" {
			matches = append(matches, match)
		}
	})

	return matches
}

func (s *Scraper) parseMatchElement(sel *goquery.Selection) models.Match {
	match := models.Match{
		Status:    "scheduled",
		StartTime: time.Now(),
	}

	teamA := sel.Find(".team-a .team-name .short-name").First()
	teamB := sel.Find(".team-b .team-name .short-name").First()

	if teamA.Length() > 0 {
		match.HomeTeam = strings.TrimSpace(teamA.Text())
	}
	if teamB.Length() > 0 {
		match.AwayTeam = strings.TrimSpace(teamB.Text())
	}

	homeScore := sel.Find(".team-a .score").First()
	awayScore := sel.Find(".team-b .score").First()

	if homeScore.Length() > 0 {
		if score, err := strconv.Atoi(strings.TrimSpace(homeScore.Text())); err == nil {
			match.HomeScore = score
		}
	}
	if awayScore.Length() > 0 {
		if score, err := strconv.Atoi(strings.TrimSpace(awayScore.Text())); err == nil {
			match.AwayScore = score
		}
	}

	gameStatus := sel.Find(".game-time").First()
	if gameStatus.Length() > 0 {
		statusText := strings.TrimSpace(gameStatus.Text())
		switch strings.ToUpper(statusText) {
		case "FT", "FULL TIME":
			match.Status = "completed"
		case "LIVE", "HT", "HALF TIME":
			match.Status = "live"
		default:
			match.Status = "scheduled"
			if parsedTime, err := s.parseTime(statusText); err == nil {
				match.StartTime = parsedTime
			}
		}
	}

	dateHeading := sel.Find("h2.date-heading").First()
	if dateHeading.Length() > 0 {
		match.Competition = strings.TrimSpace(dateHeading.Text())
	} else {
		match.Competition = "Rugby Match"
	}

	gameLink := sel.Find("a.mobileScoreboardLink").First()
	href, exists := gameLink.Attr("href")
	if exists {
		if strings.HasPrefix(href, "/") {
			match.URL = "https://www.espn.com" + href
		} else {
			match.URL = href
		}
	}

	return match
}

func (s *Scraper) parseTime(timeStr string) (time.Time, error) {
	timeStr = strings.TrimSpace(timeStr)

	formats := []string{
		"2:00 PM",
		"3:45 PM",
		"15:45",
		"2006-01-02 15:04",
		"01/02 15:04",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			if t.Year() == 0 {
				t = t.AddDate(time.Now().Year(), 0, 0)
			}
			return t, nil
		}
	}

	return time.Now(), nil
}

func (s *Scraper) StartPeriodicScraping(db *gorm.DB, intervalMinutes int) {
	ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
	defer ticker.Stop()

	s.ScrapeESPN(db)

	for {
		select {
		case <-ticker.C:
			s.ScrapeESPN(db)
		}
	}
}