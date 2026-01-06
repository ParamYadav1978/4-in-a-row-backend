# Connect 4 Backend

Go backend server for Connect Four game with WebSocket support, intelligent bot AI, and PostgreSQL database.

## Architecture

This backend handles:
- **Real-time Communication**: WebSocket connections for live gameplay
- **Game Logic**: Board state management and win detection
- **Bot AI**: Intelligent move selection algorithm
- **Database Operations**: Player statistics and leaderboard
- **Session Management**: Reconnection handling

##  Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── bot/
│   │   └── bot.go           # Bot AI logic
│   ├── db/
│   │   ├── db.go            # Database connection
│   │   └── leaderboard.go   # Leaderboard operations
│   ├── game/
│   │   └── board.go         # Game board logic
│   ├── health/
│   │   └── handler.go       # Health check endpoint
│   ├── httpapi/
│   │   └── leaderboard.go   # HTTP API handlers
│   ├── matchmaking/
│   │   └── matchmaking.go   # Player matching logic
│   ├── router/
│   │   └── router.go        # Route configuration
│   ├── session/
│   │   └── session.go       # Session management
│   └── ws/
│       ├── handler.go       # WebSocket handler
│       └── message.go       # Message types
├── go.mod                    # Go dependencies
└── go.sum                    # Dependency checksums
```

##  Getting Started

### Prerequisites

- Go 1.23 or higher
- PostgreSQL 14 or higher

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/ParamYadav1978/4-in-a-row-backend.git
cd 4-in-a-row-backend/backend
```

2. **Install dependencies**
```bash
go mod download
```

3. **Set up database**
```bash
# Create database
createdb connect4

# Run schema (connect to database and execute)
psql connect4
```
```sql
CREATE TABLE leaderboard (
    username VARCHAR(255) PRIMARY KEY,
    games_played INT DEFAULT 0,
    wins INT DEFAULT 0,
    losses INT DEFAULT 0,
    draws INT DEFAULT 0,
    bot_games INT DEFAULT 0
);
```

4. **Set environment variables** (optional for local development)
```bash
export DATABASE_URL="postgres://localhost/connect4?sslmode=disable"
export PORT="8080"
```

5. **Run the server**
```bash
go run ./cmd/server/main.go
```

Server will start on `http://localhost:8080`

## 🔌 API Reference

### WebSocket Endpoint

**`/ws`** - WebSocket connection for real-time gameplay

#### Client → Server Messages

```json
// Join game
{
  "type": "join",
  "username": "player123"
}

// Make a move
{
  "type": "move",
  "column": 3,
  "player": 1
}
```

#### Server → Client Messages

```json
// Connection established
{
  "type": "connected",
  "player": 1
}

// Waiting for opponent
{
  "type": "waiting",
  "username": "player123"
}

// Bot game started
{
  "type": "bot_start",
  "player": 1,
  "username": "player123"
}

// Board update
{
  "type": "board",
  "board": [[0,0,0,0,0,0,0], ...],
  "currentPlayer": 2
}

// Game over
{
  "type": "game_over",
  "winner": 1,
  "board": [[1,2,1,0,0,0,0], ...]
}
```

### HTTP Endpoints

#### Get Leaderboard
```http
GET /leaderboard
```

**Response:**
```json
[
  {
    "username": "player1",
    "games_played": 10,
    "wins": 7,
    "losses": 2,
    "draws": 1,
    "bot_games": 10
  }
]
```

##  Bot AI Algorithm

The bot uses a strategic decision tree:

1. **Check for Winning Move** - Scan all columns for immediate win
2. **Block Opponent** - Scan for opponent's winning moves and block
3. **Prefer Center** - Play column 3 (center) if available
4. **Strategic Placement** - Play near-center columns (2, 4, 1, 5)
5. **Fallback** - Play any available column

### Bot Logic Flow

```go
// 1. Try winning move
for each column:
    if can_win(column, bot_player):
        return column

// 2. Block opponent
for each column:
    if can_win(column, human_player):
        return column

// 3. Center preference
if column_3_available:
    return 3

// 4. Near-center strategy
for column in [2, 4, 1, 5]:
    if column_available:
        return column

// 5. Any column
for column in [0..6]:
    if column_available:
        return column
```

## 🎮 Game Logic

### Board Representation

```
Row 0: [0][0][0][0][0][0][0]  ← Top
Row 1: [0][0][0][0][0][0][0]
Row 2: [0][0][0][0][0][0][0]
Row 3: [0][0][1][0][0][0][0]
Row 4: [0][2][1][0][0][0][0]
Row 5: [2][1][1][2][0][0][0]  ← Bottom
       ↑
       Column 0
```

- `0` = Empty
- `1` = Player 1
- `2` = Player 2 (Bot)

### Win Detection

Checks 4 directions for each disc:
- **Horizontal**: Left ↔ Right
- **Vertical**: Up ↔ Down
- **Diagonal**: ↖ ↔ ↘
- **Anti-diagonal**: ↗ ↔ ↙

## 📊 Database Schema

```sql
CREATE TABLE leaderboard (
    username VARCHAR(255) PRIMARY KEY,
    games_played INT DEFAULT 0,
    wins INT DEFAULT 0,
    losses INT DEFAULT 0,
    draws INT DEFAULT 0,
    bot_games INT DEFAULT 0
);
```

### Database Operations

- `RecordResult(winner, loser, is_bot_game)` - Update stats after game
- `GetLeaderboard()` - Fetch all player statistics

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://localhost/connect4?sslmode=disable` |
| `PORT` | Server port | `8080` |

## 🚀 Deployment

### Render Deployment

1. **Build Command:**
```bash
go build -o bin/server ./cmd/server
```

2. **Start Command:**
```bash
./bin/server
```

3. **Environment Variables:**
- Set `DATABASE_URL` to Render PostgreSQL Internal URL
- `PORT` is provided automatically by Render

### Docker (Optional)

```dockerfile
FROM golang:1.23-alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server
EXPOSE 8080
CMD ["./server"]
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Test with coverage
go test -cover ./...

# Test specific package
go test ./internal/game
```

## 📈 Performance Considerations

- **WebSocket Connections**: Each player maintains a single WebSocket connection
- **Concurrency**: Go's goroutines handle multiple simultaneous games
- **Database**: Connection pooling for efficient DB operations
- **Session Cleanup**: 30-second timeout for disconnected players

## 🐛 Debugging

Enable debug logging:
```bash
# The server prints bot moves and connection events
go run ./cmd/server/main.go
# Output:
#  Connected to PostgreSQL
# Starting HTTP server on port 8080...
# BOT CHOSE COLUMN: 3
```

## Security Notes

- WebSocket origin checking disabled for development (add in production)
- Database credentials via environment variables only
- No authentication required for MVP (add for production)

##  Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


**Built with Go*
