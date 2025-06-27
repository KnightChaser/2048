package engine

// Direction represents a movie direction in the game
type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

// Game holds the state of a 2048 game
type Game struct {
	Board [GridN][GridN]int // 4 * 4 grid of tiles
	Score int               // accumulated score
}

// NewGame initializes a new game with two tiles spawned.
func NewGame() *Game {
	g := &Game{}

	SpawnTile(&g.Board)
	SpawnTile(&g.Board)
	return g
}

// Move applies a slide/merge in the given direction.
// Returns moved=true if any tile moved or merged (changed), and score gain.
func (g *Game) Move(dir Direction) (moved bool, gain int) {
	var lines [GridN][]int

	// NOTE: Extract rows or columns into lines based on direction.
	// It adjusts some logics so every move can be treated as a left move,
	switch dir {
	case Left:
		// Left: copy rows directly
		for row := range GridN {
			lines[row] = copyLine(g.Board[row][:])
		}
	case Right:
		// Right: copy rows and reverse them
		for row := range GridN {
			lines[row] = reverse(copyLine(g.Board[row][:]))
		}
	case Up:
		// Up: copy columns into rows
		for column := range GridN {
			col := make([]int, GridN)
			for row := range GridN {
				col[row] = g.Board[row][column]
			}
			lines[column] = col
		}
	case Down:
		// Down: copy columns and reverse t
		for column := range GridN {
			col := make([]int, GridN)
			for row := range GridN {
				col[row] = g.Board[GridN-1-row][column] // reverse column for down
			}
			lines[column] = col
		}
	}

	// NOTE: Process each line
	// 1. Slides all non-zero left.
	// 2. Merges identical neighbors (doubling one, zeroing the other, adding to gain).
	// 3. Slides again to collapse the gaps.
	for i, line := range lines {
		newLine, movedLine, gainLine := slideMergeLine(line)
		if movedLine {
			moved = true
		}
		gain += gainLine

		// Write back to board
		// NOTE: Take each transformed slice and shove it back into the board
		// in the correct orientation (undoing the reserce if needed).
		switch dir {
		case Left:
			// Left: write new line back to the same row
			for c, v := range newLine {
				g.Board[i][c] = v
			}
		case Right:
			// Right: write new line back to the same row, reversed
			for c, v := range newLine {
				g.Board[i][GridN-1-c] = v // reverse back to original row
			}
		case Up:
			// Up: write new line back to the same column
			for r, v := range newLine {
				g.Board[r][i] = v
			}
		case Down:
			// Down: write new line back to the same column, reversed
			for r, v := range newLine {
				g.Board[GridN-1-r][i] = v // reverse back to original column
			}
		}
	}

	if moved {
		g.Score += gain
	}

	return moved, gain
}

// copyLine clones a slice of ints.
func copyLine(line []int) []int {
	out := make([]int, len(line))
	copy(out, line)
	return out
}

// reverse reverses a slice of ints.
func reverse(line []int) []int {
	n := len(line)
	for i := range line[:n/2] {
		line[i], line[n-i-1] = line[n-i-1], line[i]
	}
	return line
}

// CanMove returns true if at least one move in possible
func (g *Game) CanMove() bool {
	// any empty cell?
	for row := range GridN {
		for column := range GridN {
			if g.Board[row][column] == 0 {
				// Empty cell is found
				return true
			}
		}
	}

	// If any adjacent cells are equal, we can merge,
	// so there will be a move possible
	for row := range GridN {
		for column := 0; column < GridN-1; column++ {
			if g.Board[row][column] == g.Board[row][column+1] {
				return true
			}
		}
	}

	for column := range GridN {
		for row := 0; row < GridN-1; row++ {
			if g.Board[row][column] == g.Board[row+1][column] {
				return true
			}
		}
	}

	return false
}
