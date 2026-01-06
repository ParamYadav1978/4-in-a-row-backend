package session

import (
	"sync"
	"time"

	"example.com/connectfour/internal/game"
)

type GameSession struct {
	Username         string
	Board            *game.Board
	CurrentPlayer    int
	GameOver         bool
	DisconnectTimer  *time.Timer
	OpponentUsername string // For PvP - stores opponent's username
}

var (
	sessions = make(map[string]*GameSession)
	mutex    sync.Mutex
)

// CreateSession creates a new session
func CreateSession(username string, board *game.Board, currentPlayer int) {
	mutex.Lock()
	defer mutex.Unlock()

	sessions[username] = &GameSession{
		Username:      username,
		Board:         board,
		CurrentPlayer: currentPlayer,
		GameOver:      false,
	}
}

// GetSession returns session if exists
func GetSession(username string) (*GameSession, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	s, ok := sessions[username]
	return s, ok
}

// RemoveSession deletes session
func RemoveSession(username string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(sessions, username)
}
