package engine

import (
	"math/rand"
)

// SpawnTile picks a random empty cell on the board and places a new tile (2 or 4)
// Returns true if a file was spawned, false if the board is full.
func SpawnTile(board *[GridN][GridN]int) bool {
	type coordinate struct {
		row    int
		column int
	}
	var empties []coordinate

	// Collect empty positions
	for row := range board {
		for column := range board[row] {
			if board[row][column] == 0 {
				empties = append(empties, coordinate{row, column})
			}
		}
	}
	if len(empties) == 0 {
		// No empty cells, can't spawn a tile
		return false
	}

	// Choose a random empty cell
	// - 90% chance of 2,
	// - 10% chance of 4
	pos := empties[rand.Intn(len(empties))]
	value := 2
	if rand.Float64() < 0.1 {
		value = 4
	}
	board[pos.row][pos.column] = value

	return true
}
