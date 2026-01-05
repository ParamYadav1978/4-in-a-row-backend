package ws
// Defines WebSocket message formats

type Message struct {
	Type          string  `json:"type"`           // move, board, error, connected
	Column        int     `json:"column"`         // column index
	Player        int     `json:"player"`         // player making the move
	Board         [][]int `json:"board"`          // board state
	CurrentPlayer int     `json:"currentPlayer"`  // whose turn is next
	Username string `json:"username"` // player's username
	Winner int `json:"winner,omitempty"`


}
