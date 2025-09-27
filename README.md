# Rugby Promiedos

A rugby scheduling backend application that provides comprehensive match information, schedules, and results by scraping data from ESPN Rugby.

## Features

- **Live Match Tracking**: Real-time updates on rugby matches
- **Match Scheduling**: View matches by date, competition, or status
- **Calendar View**: Monthly calendar with match summaries
- **Automated Scraping**: Periodic data collection from ESPN Rugby
- **RESTful API**: Clean JSON endpoints for frontend integration

## Tech Stack

- **Backend**: Go 1.21+
- **Web Framework**: Fiber
- **Database**: SQLite (with GORM ORM)
- **Web Scraping**: Goquery
- **HTTP Client**: Built-in Go HTTP client

## Quick Start

### Prerequisites

- Go 1.21 or higher
- SQLite3

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd rugby-promiedos
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Environment Variables

- `PORT`: Server port (default: 8080)
- `DATABASE_URL`: Database file path (default: rugby.db)

## API Endpoints

### Matches
- `GET /api/matches/today` - Get today's matches
- `GET /api/matches/upcoming` - Get upcoming matches (next 7 days)
- `GET /api/matches/date/:date` - Get matches for specific date (YYYY-MM-DD)
- `GET /api/matches/live` - Get currently live matches

### Competitions
- `GET /api/competitions` - Get all competitions
- `GET /api/competitions/:id/matches` - Get matches for specific competition

### Calendar
- `GET /api/calendar` - Get monthly calendar view
  - Query params: `year`, `month`

## API Response Format

All responses follow this structure:
```json
{
  "success": true,
  "message": "Success message (optional)",
  "data": {}
}
```

## Data Models

### Match
```json
{
  "id": "uuid",
  "home_team": "Team A",
  "away_team": "Team B",
  "home_score": 0,
  "away_score": 0,
  "status": "scheduled|live|completed",
  "start_time": "2024-01-01T15:00:00Z",
  "competition": "Competition Name",
  "venue": "Stadium Name",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Development

### Project Structure
```
rugby-promiedos/
├── main.go              # Application entry point
├── go.mod              # Go module file
├── internal/
│   ├── config/         # Configuration
│   ├── database/       # Database connection
│   ├── handlers/       # HTTP handlers
│   ├── models/         # Data models
│   └── scraper/        # Web scraping logic
└── PRD.md              # Product Requirements Document
```

### Running Tests
```bash
go test ./...
```

## Future Enhancements

- Frontend web application
- User authentication and favorites
- Push notifications for match starts
- Detailed match statistics
- Mobile app
- Social features and comments
- Additional data sources

## License

This project is licensed under the MIT License.