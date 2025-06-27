package engine

import (
	"testing"
)

// TestSpawnTileEmpty ensures that spawning on an empty board places exactly one tile (2 or 4).
func TestSpawnTileEmpty(t *testing.T) {
	// Initialize empty board
	board := &[GridN][GridN]int{}

	ok := SpawnTile(board)
	if !ok {
		t.Fatal("SpawnTile returned false on an empty board, expected true")
	}

	count := 0
	for r := range board {
		for c := range board[r] {
			v := board[r][c]
			if v != 0 {
				count++
				if v != 2 && v != 4 {
					t.Errorf("unexpected tile value %d; want 2 or 4", v)
				}
			}
		}
	}
	if count != 1 {
		t.Errorf("expected exactly 1 new tile, got %d", count)
	}
}

// TestSpawnTileFull ensures that spawning on a full board returns false and doesn't modify the board.
func TestSpawnTileFull(t *testing.T) {
	// Initialize full board
	board := &[GridN][GridN]int{}
	for r := range board {
		for c := range board[r] {
			board[r][c] = 2
		}
	}

	// Copy original for later comparison
	op := *board

	ok := SpawnTile(board)
	if ok {
		t.Fatal("SpawnTile returned true on a full board, expected false")
	}
	if *board != op {
		t.Error("board changed on full spawn attempt, expected no modifications")
	}
}

// TestSpawnTileSingleEmpty ensures that spawning when exactly one cell is empty fills that cell.
func TestSpawnTileSingleEmpty(t *testing.T) {
	// Initialize full board except one cell
	board := &[GridN][GridN]int{}
	for r := range board {
		for c := range board[r] {
			board[r][c] = 2
		}
	}
	// Clear one cell
	emptyR, emptyC := 1, 2
	board[emptyR][emptyC] = 0

	ok := SpawnTile(board)
	if !ok {
		t.Fatal("SpawnTile returned false when one cell was empty, expected true")
	}
	if board[emptyR][emptyC] == 0 {
		t.Errorf("expected cell [%d][%d] to be filled, but it's still zero", emptyR, emptyC)
	}
}
