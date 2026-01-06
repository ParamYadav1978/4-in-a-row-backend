import { useEffect, useRef, useState } from "react";
import Board from "./components/Board";
import Leaderboard from "./components/Leaderboard";
import UsernameModal from "./components/UsernameModal";
import "./styles/App.css";

function App() {
  const wsRef = useRef(null);
  const playerNumberRef = useRef(null); // Use ref to avoid closure issues

  const emptyBoard = Array.from({ length: 6 }, () =>
    Array(7).fill(0)
  );
  const [gameOver, setGameOver] = useState(false);

  const [board, setBoard] = useState(emptyBoard);
  const [currentPlayer, setCurrentPlayer] = useState(1);
  const [username, setUsername] = useState("");
  const [connected, setConnected] = useState(false);
  const [playerNumber, setPlayerNumber] = useState(null);
  const [leaderboardKey, setLeaderboardKey] = useState(0);
  const [showLeaderboard, setShowLeaderboard] = useState(false);
  const [playerColor, setPlayerColor] = useState("#FF6B6B"); // Default red
  const [gameStarted, setGameStarted] = useState(false); // Track if bot game started
  const [countdown, setCountdown] = useState(10); // Countdown timer


  // Connect WebSocket
  const connect = (name) => {
    const ws = new WebSocket("wss://four-in-a-row-backend-qq63.onrender.com/ws");

    ws.onopen = () => {
      ws.send(
        JSON.stringify({
          type: "join",
          username: name,
        })
      );
      setConnected(true);
    };

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);

      if (msg.type === "connected") {
        // Set player number immediately on connection
        playerNumberRef.current = msg.player || 1;
        setPlayerNumber(msg.player || 1);
      }

      if (msg.type === "waiting") {
        // Start countdown when waiting for opponent
        setGameStarted(false);
        setCountdown(10);
        const interval = setInterval(() => {
          setCountdown(prev => {
            if (prev <= 1) {
              clearInterval(interval);
              return 0;
            }
            return prev - 1;
          });
        }, 1000);
      }

      if (msg.type === "bot_start") {
        // Bot fallback triggered after 10 seconds
        playerNumberRef.current = msg.player;
        setPlayerNumber(msg.player);
        setBoard(emptyBoard);
        setCurrentPlayer(1);
        setGameOver(false);
        setGameStarted(true);
        setCountdown(0);
      }

      if (msg.type === "matched") {
        playerNumberRef.current = msg.player;
        setPlayerNumber(msg.player);
        setBoard(emptyBoard);
        setCurrentPlayer(1);
        setGameOver(false);
        setGameStarted(true);
        setCountdown(0);
      }

      if (msg.type === "board") {
        setBoard(msg.board);
        setCurrentPlayer(msg.currentPlayer);
      }

      if (msg.type === "reconnected") {
        setBoard(msg.board);
        setCurrentPlayer(msg.currentPlayer);
      }

      if (msg.type === "game_over") {
        setGameOver(true);
        setLeaderboardKey(k => k + 1);
        alert(
          msg.winner === playerNumberRef.current
            ? "🎉 You Win!"
            : "🤖 Bot Wins!"
        );
      }
    };

    ws.onclose = () => {
      setConnected(false);
    };

    wsRef.current = ws;
    setUsername(name);
    setGameOver(false);
  };

  // Send move
  const playMove = (col) => {
    if (!connected) return;
    if (!gameStarted) return; // Don't allow moves until game starts
    if (currentPlayer !== playerNumber) return;
    if (gameOver) return;

    wsRef.current.send(
      JSON.stringify({
        type: "move",
        column: col,
        player: playerNumber,
      })
    );
  };

  return (
    <div className="app-container">
      {!username && <UsernameModal onSubmit={connect} />}

      <div className="game-header">
        <h1>🎮 Connect 4</h1>
        <p className="status-text">
          {gameOver
            ? "🏁 Game Over"
            : !gameStarted && countdown > 0
            ? `⏳ Finding opponent... Bot starts in ${countdown}s`
            : !gameStarted
            ? "🤖 Starting bot game..."
            : currentPlayer === playerNumber
            ? "🎯 Your Turn"
            : "🤖 Bot Thinking..."}
        </p>
      </div>

      <div className="game-wrapper">
        <div className="left-panel">
          <div className="color-picker">
            <label>Your Disc Color:</label>
            <div className="color-options">
              {[
                { color: "#FF6B6B", name: "Red" },
                { color: "#4ECDC4", name: "Teal" },
                { color: "#45B7D1", name: "Blue" },
                { color: "#FFA07A", name: "Coral" },
                { color: "#9B59B6", name: "Purple" },
              ].map((option) => (
                <button
                  key={option.color}
                  className={`color-btn ${
                    playerColor === option.color ? "active" : ""
                  }`}
                  style={{ backgroundColor: option.color }}
                  onClick={() => setPlayerColor(option.color)}
                  title={option.name}
                  disabled={gameOver}
                ></button>
              ))}
            </div>
            <div
              className="color-preview"
              style={{ backgroundColor: playerColor }}
            ></div>
          </div>

          <button
            className="leaderboard-btn"
            onClick={() => setShowLeaderboard(!showLeaderboard)}
          >
            {showLeaderboard ? "✕ Close" : "🏆 Leaderboard"}
          </button>
        </div>

        <div className="board-container">
          <Board board={board} onPlay={playMove} playerColor={playerColor} />
        </div>
      </div>

      {showLeaderboard && (
        <div className="leaderboard-overlay" onClick={() => setShowLeaderboard(false)}>
          <div className="leaderboard-sidebar" onClick={(e) => e.stopPropagation()}>
            <div className="leaderboard-header">
              <h3>🏆 Leaderboard</h3>
              <button
                className="close-btn"
                onClick={() => setShowLeaderboard(false)}
              >
                ✕
              </button>
            </div>
            <Leaderboard refreshKey={leaderboardKey} />
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
