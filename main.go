package main

import (
	"log"
	"time"
	"rugby-promiedos/internal/config"
	"rugby-promiedos/internal/database"
	"rugby-promiedos/internal/handlers"
	"rugby-promiedos/internal/models"
	"rugby-promiedos/internal/scraper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(db, &models.Match{}, &models.Competition{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	s := scraper.NewScraper()
	go s.StartPeriodicScraping(db, 30) // Scrape every 30 minutes

	app := fiber.New(fiber.Config{
		AppName:      "Rugby Promiedos",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	handler := handlers.NewHandler(db)

	api := app.Group("/api")

	api.Get("/matches/today", handler.GetTodayMatches)
	api.Get("/matches/upcoming", handler.GetUpcomingMatches)
	api.Get("/matches/date/:date", handler.GetMatchesByDate)
	api.Get("/matches/live", handler.GetLiveMatches)
	api.Get("/competitions", handler.GetCompetitions)
	api.Get("/competitions/:id/matches", handler.GetMatchesByCompetition)
	api.Get("/calendar", handler.GetCalendar)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}