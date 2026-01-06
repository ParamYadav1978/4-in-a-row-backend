# 🎨 Connect 4 Frontend

Beautiful, animated React frontend for the Connect Four game with real-time gameplay and smooth user experience.

## ✨ Features

### Visual Features
- **Gradient Backgrounds**: Beautiful animated gradient backgrounds
- **Smooth Animations**: Disc drop animations with rotation and bounce effects
- **Fade-in Effects**: Elegant component transitions
- **Responsive Design**: Works seamlessly on all screen sizes
- **Custom Scrollbar**: Styled scrollbar for leaderboard sidebar

### Game Features
- **Color Picker**: Choose from 5 vibrant disc colors
  - Red (#FF6B6B)
  - Teal (#26A69A)
  - Blue (#42A5F5)
  - Coral (#FF7043)
  - Purple (#AB47BC)
- **Countdown Timer**: Visual countdown showing bot start time
- **Toggleable Leaderboard**: Slide-in sidebar with player statistics
- **Game Status**: Clear indicators for turn status
- **Username Modal**: Clean username entry interface

### Technical Features
- **WebSocket Connection**: Real-time bidirectional communication
- **React State Management**: Efficient state updates with refs
- **Component Architecture**: Modular, reusable components
- **CSS Animations**: Performant keyframe animations

## 📁 Project Structure

```
connect4-frontend/
├── src/
│   ├── components/
│   │   ├── Board.jsx           # Game board component
│   │   ├── Cell.jsx            # Individual cell component
│   │   ├── Leaderboard.jsx     # Leaderboard display
│   │   └── UsernameModal.jsx   # Username input modal
│   ├── styles/
│   │   ├── App.css             # Main app styles
│   │   └── board.css           # Board-specific styles
│   ├── App.jsx                 # Main app component
│   └── main.jsx                # Application entry point
├── public/                      # Static assets
├── index.html                   # HTML template
├── package.json                 # Dependencies
└── vite.config.js              # Vite configuration
```

## 🚀 Getting Started

### Prerequisites

- Node.js 18 or higher
- npm or yarn

### Installation

1. **Navigate to frontend directory**
```bash
cd frontend/connect4-frontend
```

2. **Install dependencies**
```bash
npm install
```

3. **Start development server**
```bash
npm run dev
```

Application will start on `http://localhost:5173`

4. **Build for production**
```bash
npm run build
```

## 🎮 Components

### App.jsx
Main application component managing:
- WebSocket connection
- Game state (board, current player, game over)
- Player state (username, player number, color)
- Leaderboard visibility
- Countdown timer

```jsx
const [board, setBoard] = useState(emptyBoard);
const [currentPlayer, setCurrentPlayer] = useState(1);
const [playerColor, setPlayerColor] = useState("#FF6B6B");
const [gameStarted, setGameStarted] = useState(false);
const [countdown, setCountdown] = useState(10);
```

### Board.jsx
Renders the 6x7 game board:
- 7 columns
- 6 rows
- Click handlers for column selection
- Passes player color to cells

### Cell.jsx
Individual cell component:
- Displays empty, player 1, or player 2 disc
- Applies custom player color
- Handles drop animations

### Leaderboard.jsx
Displays player statistics:
- Fetches data from backend API
- Sorts by wins (descending)
- Shows username, games played, wins, losses, draws

### UsernameModal.jsx
Username entry interface:
- Modal overlay
- Input field with validation
- Submit button

## 🎨 Styling

### Color Scheme

```css
/* Available disc colors */
--red: #FF6B6B;
--teal: #26A69A;
--blue: #42A5F5;
--coral: #FF7043;
--purple: #AB47BC;

/* Background gradients */
background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
```

### Animations

**Fade In:**
```css
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
```

**Slide Down:**
```css
@keyframes slideDown {
  from { transform: translateY(-20px); opacity: 0; }
  to { transform: translateY(0); opacity: 1; }
}
```

**Drop Bounce:**
```css
@keyframes dropBounce {
  0% { transform: translateY(-500px) rotate(0deg); }
  60% { transform: translateY(10px) rotate(360deg); }
  80% { transform: translateY(-5px); }
  100% { transform: translateY(0); }
}
```

**Slide In From Right:**
```css
@keyframes slideInFromRight {
  from { transform: translateX(100%); }
  to { transform: translateX(0); }
}
```

## 🔌 WebSocket Integration

### Connection Setup

```javascript
const ws = new WebSocket("wss://four-in-a-row-backend-qq63.onrender.com/ws");

ws.onopen = () => {
  ws.send(JSON.stringify({
    type: "join",
    username: name
  }));
};
```

### Message Handlers

```javascript
ws.onmessage = (event) => {
  const msg = JSON.parse(event.data);
  
  // Handle different message types
  switch(msg.type) {
    case "connected": // Initial connection
    case "waiting": // Matchmaking
    case "bot_start": // Bot game started
    case "board": // Board update
    case "game_over": // Game finished
  }
};
```

### Sending Moves

```javascript
const playMove = (col) => {
  wsRef.current.send(JSON.stringify({
    type: "move",
    column: col,
    player: playerNumber
  }));
};
```

## 🎯 State Management

### Using Refs for WebSocket Closure Issues

```javascript
const playerNumberRef = useRef(null);

// Update both state and ref
setPlayerNumber(value);
playerNumberRef.current = value;

// Use ref in WebSocket callbacks
msg.winner === playerNumberRef.current
```

### Countdown Timer

```javascript
const interval = setInterval(() => {
  setCountdown(prev => {
    if (prev <= 1) {
      clearInterval(interval);
      return 0;
    }
    return prev - 1;
  });
}, 1000);
```

## 🚀 Deployment

### Vercel Deployment

1. **Install Vercel CLI**
```bash
npm install -g vercel
```

2. **Deploy**
```bash
vercel
```

3. **Configure**
- Build Command: `npm run build`
- Output Directory: `dist`
- Install Command: `npm install`

### Environment Variables

For production, create `.env`:
```
VITE_API_URL=https://four-in-a-row-backend-qq63.onrender.com
```

Update WebSocket URL:
```javascript
const ws = new WebSocket(import.meta.env.VITE_API_URL + "/ws");
```

## 📱 Responsive Design

### Breakpoints

```css
/* Mobile */
@media (max-width: 768px) {
  .board { scale: 0.8; }
  .leaderboard-sidebar { width: 100%; }
}

/* Tablet */
@media (min-width: 769px) and (max-width: 1024px) {
  .board { scale: 0.9; }
}

/* Desktop */
@media (min-width: 1025px) {
  .board { scale: 1; }
}
```

## 🧪 Testing

```bash
# Run linter
npm run lint

# Fix linting issues
npm run lint:fix

# Preview production build
npm run preview
```

## ⚡ Performance Optimization

- **Code Splitting**: Dynamic imports for heavy components
- **Memoization**: React.memo for Cell components
- **Lazy Loading**: Defer leaderboard data fetch
- **CSS Animations**: Use transform and opacity for GPU acceleration
- **WebSocket**: Single persistent connection

## 🎨 Customization

### Add New Disc Color

1. **Add color to App.jsx:**
```javascript
const colors = [
  { name: "Red", value: "#FF6B6B" },
  { name: "Green", value: "#66BB6A" }, // New color
  // ...
];
```

2. **Update color picker UI** in render section

### Change Animation Speed

```css
/* In board.css */
.cell.player1.drop {
  animation: dropBounce 0.6s ease-out; /* Adjust duration */
}
```

## 🐛 Troubleshooting

### WebSocket Connection Failed
```
Error: WebSocket connection to '...' failed
```
**Solution:** Check backend is running and URL is correct

### Board Not Updating
**Solution:** Ensure player number is set before making moves

### Animation Not Working
**Solution:** Check CSS is imported and class names match

## 🤝 Contributing

1. Follow React best practices
2. Use functional components with hooks
3. Keep components small and focused
4. Add PropTypes for type checking
5. Write clear comments for complex logic

## 📝 License

MIT License - see LICENSE file for details

---

**Built with React ⚛️ and Vite ⚡**
