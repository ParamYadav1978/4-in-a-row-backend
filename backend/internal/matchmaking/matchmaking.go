package matchmaking
// Simple waiting room for players

import (
	"sync"
	"time"
	"github.com/gorilla/websocket"
)


// Player represents a user waiting to be matched
type Player struct {
	Username string
	IsBot    bool
	Conn     *websocket.Conn
}

// waitingPlayer holds the currently waiting player
var waitingPlayer *Player

// mutex protects access to waitingPlayer
var mutex sync.Mutex

var waitingTimer *time.Timer

// AddPlayer adds a player to the waiting room
// If another player is already waiting, it returns that player (match found)
// If not, it stores the player and returns nil
// AddPlayer adds a player to the waiting room with timeout logic
func AddPlayer(player *Player, onTimeout func(*Player)) *Player {

	mutex.Lock()
	defer mutex.Unlock()

	// If no one is waiting, store player and start timer
	if waitingPlayer == nil {

		waitingPlayer = player

		// Start 10-second timer
		waitingTimer = time.AfterFunc(10*time.Second, func() {

			mutex.Lock()
			defer mutex.Unlock()

			// If player still waiting, trigger bot fallback
			if waitingPlayer != nil && waitingPlayer.Username == player.Username {
				waitingPlayer = nil
				onTimeout(player)
			}
		})

		return nil
	}

	// Someone waiting → cancel timer and match
	if waitingTimer != nil {
		waitingTimer.Stop()
	}

	matchedPlayer := waitingPlayer
	waitingPlayer = nil
	return matchedPlayer
}
