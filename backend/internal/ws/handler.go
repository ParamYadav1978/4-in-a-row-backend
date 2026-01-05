package ws

// Handles WebSocket connections for the game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/connectfour/internal/bot"
	"example.com/connectfour/internal/db"
	"example.com/connectfour/internal/game"
	"example.com/connectfour/internal/matchmaking"
	"example.com/connectfour/internal/session"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader (HTTP → WS)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleGameSocket handles WebSocket connections
func HandleGameSocket(w http.ResponseWriter, r *http.Request) {

	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// Create a new game board for this connection
	board := game.NewBoard()

	// Player 1 always starts
	currentPlayer := game.Player1

	isBotGame := false
	botPlayer := game.Player2
	humanPlayer := game.Player1
	gameOver := false

	var username string // stores username for this connection

	// Send connected message (JSON only)
	connectedMsg := Message{
		Type: "connected",
	}
	response, _ := json.Marshal(connectedMsg)
	conn.WriteMessage(websocket.TextMessage, response)

	// Listen for messages from client
	for {

		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected:", username)

			sess, ok := session.GetSession(username)
			if ok {
				sess.DisconnectTimer = time.AfterFunc(30*time.Second, func() {
					fmt.Println("Player did not reconnect:", username)
					sess.GameOver = true
					session.RemoveSession(username)
				})
			}

			break
		}

		// Parse JSON
		var msg Message
		err = json.Unmarshal(messageBytes, &msg)
		if err != nil {
			continue
		}

		switch msg.Type {

		case "join":

			// Save username
			username = msg.Username

			// Create / reset session
			session.CreateSession(username, board, currentPlayer)

			player := &matchmaking.Player{
				Username: username,
				Conn:     conn,
			}

			// Try to find opponent
			opponent := matchmaking.AddPlayer(player, func(p *matchmaking.Player) {
				// Bot fallback triggered after 10 seconds
				isBotGame = true
				botMsg := Message{
					Type:     "bot_start",
					Username: p.Username,
				}
				response, _ := json.Marshal(botMsg)
				p.Conn.WriteMessage(websocket.TextMessage, response)
			})

			if opponent == nil {
				// No opponent found - waiting for second player or bot timeout
				waitMsg := Message{
					Type:     "waiting",
					Username: username,
				}
				response, _ := json.Marshal(waitMsg)
				conn.WriteMessage(websocket.TextMessage, response)
				continue
			}

			// Opponent found - PvP match
			isBotGame = false
			gameOver = false
			currentPlayer = game.Player1

			// Send matched to both players
			opponent.Conn.WriteJSON(Message{
				Type:     "matched",
				Username: username,
			})

			conn.WriteJSON(Message{
				Type:     "matched",
				Username: opponent.Username,
			})
			continue

		case "move":

			if gameOver {
				return
			}

			// Reject move if wrong player tries to play
			if msg.Player != currentPlayer {
				errorMsg := Message{
					Type: "error",
				}
				response, _ := json.Marshal(errorMsg)
				conn.WriteMessage(websocket.TextMessage, response)
				continue
			}

			// Apply move for current player
			success := board.DropDisc(msg.Column, currentPlayer)

			if !success {
				// Invalid move
				errorMsg := Message{Type: "error"}
				response, _ := json.Marshal(errorMsg)
				conn.WriteMessage(websocket.TextMessage, response)
				continue
			}

			if board.CheckWin(currentPlayer) {
				gameOver = true

				// 🏆 UPDATE LEADERBOARD (bot games only for now)
				if isBotGame {
					if currentPlayer == game.Player1 {
						db.RecordResult(username, "BOT", true)
					} else {
						db.RecordResult("BOT", username, true)
					}
				}

				winMsg := Message{
					Type:   "game_over",
					Winner: currentPlayer,
					Board:  board.Grid,
				}

				response, _ := json.Marshal(winMsg)
				conn.WriteMessage(websocket.TextMessage, response)
				return
			}

			// Switch turn after successful move
			if currentPlayer == game.Player1 {
				currentPlayer = game.Player2
			} else {
				currentPlayer = game.Player1
			}

			// ✅ SAVE SESSION HERE
			session.CreateSession(username, board, currentPlayer)
			// Send updated board
			boardMsg := Message{
				Type:          "board",
				Board:         board.Grid,
				CurrentPlayer: currentPlayer,
			}
			response, _ = json.Marshal(boardMsg)
			conn.WriteMessage(websocket.TextMessage, response)

			// 🤖 BOT MOVE (if bot game and bot's turn)
			if isBotGame && currentPlayer == botPlayer {
				// Add 1 second delay to make bot feel more natural
				time.Sleep(1 * time.Second)

				botCol := bot.ChooseMove(board, botPlayer, humanPlayer)
				fmt.Println("BOT CHOSE COLUMN:", botCol)

				if botCol != -1 {

					board.DropDisc(botCol, botPlayer)

					if board.CheckWin(botPlayer) {
						gameOver = true

						// 🏆 UPDATE LEADERBOARD
						db.RecordResult("BOT", username, true)

						winMsg := Message{
							Type:   "game_over",
							Winner: botPlayer,
							Board:  board.Grid,
						}

						response, _ := json.Marshal(winMsg)
						conn.WriteMessage(websocket.TextMessage, response)
						return
					}

					// switch back to human ONLY if bot didn't win
					currentPlayer = humanPlayer
					// ✅ SAVE SESSION HERE
					session.CreateSession(username, board, currentPlayer)

					botBoardMsg := Message{
						Type:          "board",
						Board:         board.Grid,
						CurrentPlayer: currentPlayer,
						Player:        botPlayer,
						Column:        botCol,
					}

					response, _ := json.Marshal(botBoardMsg)
					conn.WriteMessage(websocket.TextMessage, response)
				}

			}

		default:
			// Unsupported message type
			errorMsg := Message{Type: "error"}
			response, _ := json.Marshal(errorMsg)
			conn.WriteMessage(websocket.TextMessage, response)
		}
	}
}
