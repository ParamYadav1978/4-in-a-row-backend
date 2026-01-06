package bot
// Competitive bot decision logic

import "example.com/connectfour/internal/game"

// isColumnPlayable checks if column has space
func isColumnPlayable(board *game.Board, col int) bool {
	return board.Grid[0][col] == 0
}


// ChooseMove returns the column the bot should play
func ChooseMove(board *game.Board, botPlayer int, humanPlayer int) int {

	// 1️⃣ Try winning move
	for col := 0; col < 7; col++ {
		copyBoard := board.Copy()
		if copyBoard.DropDisc(col, botPlayer) && copyBoard.CheckWin(botPlayer) {
			return col
		}
	}

	// 2️⃣ Block opponent winning move
	for col := 0; col < 7; col++ {
		copyBoard := board.Copy()
		if copyBoard.DropDisc(col, humanPlayer) && copyBoard.CheckWin(humanPlayer) {
			return col
		}
	}

	// 3️⃣ Center column priority
	if isColumnPlayable(board, 3) {
		return 3
	}

	// 4️⃣ Prefer near-center columns
	preferred := []int{2, 4, 1, 5}
	for _, col := range preferred {
		if isColumnPlayable(board, col) {
			return col
		}
	}

	// 5️⃣ Final fallback: any playable column
	for col := 0; col < 7; col++ {
		if isColumnPlayable(board, col) {
			return col
		}
	}

	return -1
}
