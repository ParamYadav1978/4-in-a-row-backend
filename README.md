# 🎮 Connect 4 - Full Stack Real-Time Game

A modern, real-time implementation of the classic Connect Four game with a beautiful animated interface, intelligent bot opponent, and competitive leaderboard system.

[![Live Demo](https://img.shields.io/badge/demo-live-success)](https://4-in-a-row-backend.vercel.app)
[![GitHub](https://img.shields.io/badge/github-repository-blue)](https://github.com/ParamYadav1978/4-in-a-row-backend)

## 🌟 Features

### Game Features
- **Real-time Gameplay**: Instant move updates using WebSocket connections
- **Smart Bot Opponent**: Competitive AI that blocks your wins and seeks its own
- **10-Second Matchmaking**: Automatic bot game after countdown if no opponent found
- **Custom Disc Colors**: Choose from 5 vibrant colors (Red, Teal, Blue, Coral, Purple)
- **Smooth Animations**: Disc drop animations with rotation and bounce effects
- **Win Detection**: Automatic detection of horizontal, vertical, and diagonal wins

### UI/UX Features
- **Animated Interface**: Gradient backgrounds with fade-in and slide-down effects
- **Countdown Timer**: Visual countdown showing "Bot starts in Xs"
- **Toggleable Leaderboard**: Slide-in sidebar with player statistics
- **Responsive Design**: Works seamlessly on desktop and mobile devices
- **Game Status Indicators**: Clear visual feedback for turn status and game state

### Backend Features
- **WebSocket Server**: Real-time bidirectional communication
- **PostgreSQL Database**: Persistent storage for player statistics
- **Bot Logic**: Intelligent move selection with win/block detection
- **Session Management**: 30-second reconnection window for disconnected players
- **Leaderboard System**: Track wins, losses, draws, and games played

## 🚀 Live Demo

**Play Now**: [https://4-in-a-row-backend.vercel.app](https://4-in-a-row-backend.vercel.app)

Simply enter your username and start playing! The game will automatically match you with a bot opponent after 10 seconds.

## 🛠️ Tech Stack

### Frontend
- **React 18** - UI framework
- **Vite** - Build tool and dev server
- **WebSocket API** - Real-time communication
- **CSS3** - Animations and styling

### Backend
- **Go** - High-performance backend server
- **Gorilla WebSocket** - WebSocket implementation
- **PostgreSQL** - Relational database
- **lib/pq** - PostgreSQL driver for Go

### Deployment
- **Vercel** - Frontend hosting
- **Render** - Backend and database hosting

## 📁 Project Structure

```
4-in-a-row-backend/
├── backend/                 # Go backend server
│   ├── cmd/
│   │   └── server/         # Main server entry point
│   ├── internal/
│   │   ├── bot/           # Bot AI logic
│   │   ├── db/            # Database operations
│   │   ├── game/          # Game board logic
│   │   ├── ws/            # WebSocket handlers
│   │   ├── httpapi/       # HTTP endpoints
│   │   ├── router/        # Route configuration
│   │   ├── session/       # Session management
│   │   └── matchmaking/   # Player matching logic
│   └── go.mod
│
└── frontend/
    └── connect4-frontend/  # React frontend
        ├── src/
        │   ├── components/  # React components
        │   ├── styles/      # CSS stylesheets
        │   ├── App.jsx      # Main app component
        │   └── main.jsx     # Entry point
        └── package.json
```

## 🎯 How to Play

1. **Enter Username**: Type your name to join the game
2. **Wait for Match**: A 10-second countdown will start looking for opponents
3. **Choose Color**: Select your disc color from 5 available options
4. **Make Moves**: Click on any column to drop your disc
5. **Win the Game**: Connect 4 discs horizontally, vertically, or diagonally
6. **Check Leaderboard**: Toggle the sidebar to see player rankings

## 💻 Local Development

### Prerequisites
- **Go 1.23+** - [Download](https://golang.org/dl/)
- **Node.js 18+** - [Download](https://nodejs.org/)
- **PostgreSQL 14+** - [Download](https://www.postgresql.org/download/)

### Backend Setup

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod download

# Set up database
createdb connect4

# Set environment variables (optional for local)
export DATABASE_URL="postgres://localhost/connect4?sslmode=disable"
export PORT="8080"

# Run the server
go run ./cmd/server/main.go
```

The backend will start on `http://localhost:8080`

### Frontend Setup

```bash
# Navigate to frontend directory
cd frontend/connect4-frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will start on `http://localhost:5173`

## 🗄️ Database Schema

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

## 🔌 API Endpoints

### WebSocket
- **`/ws`** - WebSocket connection for real-time gameplay

### HTTP
- **`GET /leaderboard`** - Fetch leaderboard data

## 🎨 Game Design Decisions

### Why Bot-Only Mode?
Currently, the game focuses on single-player vs bot gameplay for stability and immediate playability. PvP mode requires additional complexity for:
- Shared game state management
- Player synchronization
- Concurrent move validation
- Disconnection handling for both players

### Bot Intelligence
The bot uses a strategic algorithm:
1. **Win Detection**: Looks for immediate winning moves
2. **Block Detection**: Blocks opponent's winning moves
3. **Center Priority**: Prefers central columns for better positioning
4. **Strategic Placement**: Falls back to near-center columns

### 1-Second Bot Delay
Adds human-like thinking time for better UX - prevents instant bot moves that feel unnatural.

## 🚀 Deployment

### Frontend (Vercel)
```bash
# Frontend auto-deploys on push to main branch
# Configured via vercel.json
```

### Backend (Render)
```bash
# Backend auto-deploys on push to main branch
# Environment variables configured in Render dashboard:
# - DATABASE_URL: PostgreSQL connection string
# - PORT: Server port (provided by Render)
```

## 📊 Features Implemented

- ✅ Real-time gameplay with WebSocket
- ✅ Intelligent bot opponent
- ✅ Matchmaking with countdown timer
- ✅ Win/loss/draw detection
- ✅ Persistent leaderboard
- ✅ Session management
- ✅ Custom disc colors
- ✅ Smooth animations
- ✅ Responsive design
- ✅ Production deployment

## 🔮 Future Enhancements

- [ ] Player vs Player (PvP) mode
- [ ] Multiple difficulty levels for bot
- [ ] Game replay system
- [ ] Chat functionality
- [ ] Tournament mode
- [ ] Player profiles with avatars
- [ ] Sound effects
- [ ] Mobile app version

## 👨‍💻 Author

**Param Yadav**
- GitHub: [@ParamYadav1978](https://github.com/ParamYadav1978)
- Email: [Your Email]

## 📄 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

- Classic Connect Four game inspiration
- React and Go communities
- Vercel and Render for hosting

---

**Made with ❤️ and ☕ by Param Yadav**

---

## Game State Management
- Active games stored in memory
- Completed games stored in PostgreSQL
- Game analytics emitted asynchronously

---

## Status
Day 1 completed: System understanding and design finalized.
