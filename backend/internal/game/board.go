package game

// This file belongs to the "game" package
// Everything related to game logic will live here

// These constants represent the state of a cell in the board
// 0 means empty cell
// 1 means Player 1's disc
// 2 means Player 2's disc
const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
)

// Board is a structure that represents the Connect-4 board
// Grid is a 2D slice (like a 2D array in JavaScript).
// It will be 6 rows x 7 columns
type Board struct {
	Grid [][]int
}







// NewBoard creates and returns a new Connect-4 board
// It initializes a 6x7 grid and fills every cell with Empty
func NewBoard() *Board {

	// Create an empty slice that will hold 6 rows
	grid := make([][]int, 6)

	// Loop through each row
	for row := 0; row < 6; row++ {

		// For each row, create 7 columns
		grid[row] = make([]int, 7)

		// Fill each column with Empty value
		for col := 0; col < 7; col++ {
			grid[row][col] = Empty
		}
	}

	// Return a pointer to a Board containing the grid
	return &Board{
		Grid: grid,
	}
}






// DropDisc drops a disc into the given column for a player
// column: column index (0 to 6)
// player: Player1 or Player2
// Returns true if the move was successful, false otherwise
func (b *Board) DropDisc(column int, player int) bool {

	// Check if column index is valid
	if column < 0 || column >= 7 {
		return false
	}

	// Start from the bottom row (row 5) and go upwards
	for row := 5; row >= 0; row-- {

		// If the cell is empty, place the disc here
		if b.Grid[row][column] == Empty {
			b.Grid[row][column] = player
			return true
		}
	}

	// If we reach here, the column is full
	return false
}






// CheckVerticalWin checks if the given player has a vertical win
// player: Player1 or Player2
// Returns true if player has 4 discs vertically
func (b *Board) CheckVerticalWin(player int) bool {

	// Loop through each column (0 to 6)
	for col := 0; col < 7; col++ {

		// Counter for consecutive discs
		count := 0

		// Loop through each row from top to bottom
		for row := 0; row < 6; row++ {

			// If current cell belongs to the player
			if b.Grid[row][col] == player {
				count++

				// If 4 consecutive discs found, player wins
				if count == 4 {
					return true
				}
			} else {
				// Reset count if sequence breaks
				count = 0
			}
		}
	}

	// No vertical win found
	return false
}






// CheckHorizontalWin checks if the given player has a horizontal win
// player: Player1 or Player2
// Returns true if player has 4 discs horizontally
func (b *Board) CheckHorizontalWin(player int) bool {

	// Loop through each row (0 to 5)
	for row := 0; row < 6; row++ {

		// Counter for consecutive discs
		count := 0

		// Loop through each column from left to right
		for col := 0; col < 7; col++ {

			// If current cell belongs to the player
			if b.Grid[row][col] == player {
				count++

				// If 4 consecutive discs found, player wins
				if count == 4 {
					return true
				}
			} else {
				// Reset count if sequence breaks
				count = 0
			}
		}
	}

	// No horizontal win found
	return false
}





// CheckDiagonalWin checks if the given player has a diagonal win
// It checks both "\" and "/" diagonals
func (b *Board) CheckDiagonalWin(player int) bool {

	// ----------- Check "\" diagonals (top-left to bottom-right) -----------

	// Rows 0 to 2 (so we have space for 4 downwards)
	for row := 0; row <= 2; row++ {

		// Columns 0 to 3 (so we have space for 4 to the right)
		for col := 0; col <= 3; col++ {

			// Check 4 diagonal cells
			if b.Grid[row][col] == player &&
				b.Grid[row+1][col+1] == player &&
				b.Grid[row+2][col+2] == player &&
				b.Grid[row+3][col+3] == player {
				return true
			}
		}
	}

	// ----------- Check "/" diagonals (bottom-left to top-right) -----------

	// Rows 3 to 5 (so we can go upwards)
	for row := 3; row < 6; row++ {

		// Columns 0 to 3 (space to the right)
		for col := 0; col <= 3; col++ {

			// Check 4 diagonal cells
			if b.Grid[row][col] == player &&
				b.Grid[row-1][col+1] == player &&
				b.Grid[row-2][col+2] == player &&
				b.Grid[row-3][col+3] == player {
				return true
			}
		}
	}

	// No diagonal win found
	return false
}




// CheckWin checks if the given player has won the game
// It combines vertical, horizontal, and diagonal checks
func (b *Board) CheckWin(player int) bool {

	// Check vertical win
	if b.CheckVerticalWin(player) {
		return true
	}

	// Check horizontal win
	if b.CheckHorizontalWin(player) {
		return true
	}

	// Check diagonal win
	if b.CheckDiagonalWin(player) {
		return true
	}

	// No win found
	return false
}



// Copy creates a deep copy of the board (used for bot simulation)
func (b *Board) Copy() *Board {
	newGrid := make([][]int, len(b.Grid))

	for i := range b.Grid {
		newGrid[i] = make([]int, len(b.Grid[i]))
		copy(newGrid[i], b.Grid[i])
	}

	return &Board{
		Grid: newGrid,
	}
}

// UndoDrop removes the top-most disc from a column (used for bot simulation)
func (b *Board) UndoDrop(col int) {
	for row := 0; row < len(b.Grid); row++ {
		if b.Grid[row][col] != 0 {
			b.Grid[row][col] = 0
			return
		}
	}
}
