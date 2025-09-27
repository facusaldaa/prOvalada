# Rugby Promiedos - Product Requirements Document

## Project Overview
A rugby scheduling application similar to Promiedos, focused on providing comprehensive match information, schedules, and results for rugby competitions worldwide.

## Phase 1: Backend Foundation (MVP)

### Core Features
1. **Match Scheduling System**
   - Display all rugby matches by date
   - Show today's matches prominently
   - Display upcoming matches for the next 7 days
   - Historical match results

2. **Match Information**
   - Teams playing
   - Match time and date
   - Competition/League
   - Match status (scheduled, live, completed)
   - Score for completed matches
   - Live match updates

3. **Data Source**
   - Web scraping from ESPN Rugby (https://www.espn.com/rugby/)
   - Automated data updates
   - Data caching for performance

### Technical Requirements

#### Backend Stack
- **Language**: Go
- **Web Framework**: Fiber
- **Web Scraping**: Goquery
- **Database**: SQLite (for MVP, can scale to PostgreSQL)
- **Caching**: Redis (optional for Phase 1)

#### API Endpoints
```
GET /api/matches/today          - Get today's matches
GET /api/matches/upcoming       - Get upcoming matches (next 7 days)
GET /api/matches/date/:date     - Get matches for specific date
GET /api/matches/live           - Get currently live matches
GET /api/competitions           - Get all competitions
GET /api/competitions/:id/matches - Get matches for specific competition
```

#### Data Models
```go
type Match struct {
    ID           string    `json:"id"`
    HomeTeam     string    `json:"home_team"`
    AwayTeam     string    `json:"away_team"`
    HomeScore    int       `json:"home_score"`
    AwayScore    int       `json:"away_score"`
    Status       string    `json:"status"` // scheduled, live, completed
    StartTime    time.Time `json:"start_time"`
    Competition  string    `json:"competition"`
    Venue        string    `json:"venue"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type Competition struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Country     string `json:"country"`
    Season      string `json:"season"`
}
```

### Success Criteria
- Successfully scrape and store match data from ESPN
- API response time < 500ms
- Data updated at least every 30 minutes
- Handle at least 1000 concurrent requests
- 99% uptime for data availability

### Future Phases
- Frontend application (React/Vue)
- User authentication and favorites
- Push notifications for match starts
- Detailed match statistics
- Mobile app
- Social features and comments