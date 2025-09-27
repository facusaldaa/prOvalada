package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"rugby-promiedos/internal/models"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

type MatchResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *Handler) GetTodayMatches(c *fiber.Ctx) error {
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	var matches []models.Match
	result := h.db.Where("start_time >= ? AND start_time < ?", today, tomorrow).
		Order("start_time ASC").
		Find(&matches)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch today's matches",
		})
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    matches,
	})
}

func (h *Handler) GetUpcomingMatches(c *fiber.Ctx) error {
	now := time.Now()
	sevenDaysLater := now.Add(7 * 24 * time.Hour)

	var matches []models.Match
	result := h.db.Where("start_time >= ? AND start_time <= ?", now, sevenDaysLater).
		Where("status = ?", "scheduled").
		Order("start_time ASC").
		Find(&matches)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch upcoming matches",
		})
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    matches,
	})
}

func (h *Handler) GetMatchesByDate(c *fiber.Ctx) error {
	dateStr := c.Params("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(MatchResponse{
			Success: false,
			Message: "Invalid date format. Use YYYY-MM-DD",
		})
	}

	startOfDay := date.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	var matches []models.Match
	result := h.db.Where("start_time >= ? AND start_time < ?", startOfDay, endOfDay).
		Order("start_time ASC").
		Find(&matches)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch matches for date",
		})
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    matches,
	})
}

func (h *Handler) GetLiveMatches(c *fiber.Ctx) error {
	var matches []models.Match
	result := h.db.Where("status = ?", "live").
		Order("start_time ASC").
		Find(&matches)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch live matches",
		})
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    matches,
	})
}

func (h *Handler) GetCompetitions(c *fiber.Ctx) error {
	var competitions []models.Competition
	result := h.db.Find(&competitions)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch competitions",
		})
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    competitions,
	})
}

func (h *Handler) GetMatchesByCompetition(c *fiber.Ctx) error {
	competitionID := c.Params("id")

	var competition models.Competition
	result := h.db.Preload("Matches").First(&competition, "id = ?", competitionID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(MatchResponse{
				Success: false,
				Message: "Competition not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch competition",
		})
	}

	limitStr := c.Query("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 50
	}

	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	var matches []models.Match
	result = h.db.Where("competition_id = ?", competitionID).
		Order("start_time DESC").
		Limit(limit).
		Offset(offset).
		Find(&matches)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch matches for competition",
		})
	}

	response := map[string]interface{}{
		"competition": competition,
		"matches":     matches,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
		},
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    response,
	})
}

func (h *Handler) GetCalendar(c *fiber.Ctx) error {
	yearStr := c.Query("year", strconv.Itoa(time.Now().Year()))
	monthStr := c.Query("month", strconv.Itoa(int(time.Now().Month())))

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(MatchResponse{
			Success: false,
			Message: "Invalid year parameter",
		})
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return c.Status(http.StatusBadRequest).JSON(MatchResponse{
			Success: false,
			Message: "Invalid month parameter (1-12)",
		})
	}

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	var matches []models.Match
	result := h.db.Where("start_time >= ? AND start_time <= ?", startOfMonth, endOfMonth).
		Order("start_time ASC").
		Find(&matches)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(MatchResponse{
			Success: false,
			Message: "Failed to fetch calendar data",
		})
	}

	calendarData := h.groupMatchesByDate(matches)

	response := map[string]interface{}{
		"year":     year,
		"month":    month,
		"calendar": calendarData,
	}

	return c.JSON(MatchResponse{
		Success: true,
		Data:    response,
	})
}

func (h *Handler) groupMatchesByDate(matches []models.Match) map[string]interface{} {
	calendar := make(map[string][]models.Match)

	for _, match := range matches {
		dateKey := match.StartTime.Format("2006-01-02")
		calendar[dateKey] = append(calendar[dateKey], match)
	}

	summary := make(map[string]interface{})
	for date, matches := range calendar {
		summary[date] = map[string]interface{}{
			"total_matches": len(matches),
			"live_matches":   h.countMatchesByStatus(matches, "live"),
			"completed":      h.countMatchesByStatus(matches, "completed"),
			"scheduled":      h.countMatchesByStatus(matches, "scheduled"),
		}
	}

	return map[string]interface{}{
		"daily_matches": calendar,
		"summary":       summary,
	}
}

func (h *Handler) countMatchesByStatus(matches []models.Match, status string) int {
	count := 0
	for _, match := range matches {
		if match.Status == status {
			count++
		}
	}
	return count
}