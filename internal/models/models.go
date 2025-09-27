package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Match struct {
	ID           string    `json:"id" gorm:"primarykey"`
	HomeTeam     string    `json:"home_team" gorm:"not null"`
	AwayTeam     string    `json:"away_team" gorm:"not null"`
	HomeScore    int       `json:"home_score" gorm:"default:0"`
	AwayScore    int       `json:"away_score" gorm:"default:0"`
	Status       string    `json:"status" gorm:"not null"` // scheduled, live, completed
	StartTime    time.Time `json:"start_time" gorm:"not null"`
	Competition  string    `json:"competition" gorm:"not null"`
	Venue        string    `json:"venue"`
	CompetitionID string    `json:"competition_id"`
	URL          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Competition struct {
	ID       string    `json:"id" gorm:"primarykey"`
	Name     string    `json:"name" gorm:"not null"`
	Country  string    `json:"country"`
	Season   string    `json:"season"`
	URL      string    `json:"url"`
	Matches  []Match   `json:"matches" gorm:"foreignKey:CompetitionID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Match) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}

func (c *Competition) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}