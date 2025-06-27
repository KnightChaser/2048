package engine

// slideLIne shifts all non-zero tiles to the front
// Returns the new line and true if any tile moved.
// e.g.) [2, 0, 2, 4] -> [2, 2, 4, 0] (true)
func slideLine(line []int) ([]int, bool) {
	n := len(line)
	out := make([]int, n)
	wi := 0 // write index
	moved := false

	for ri, v := range line {
		if v == 0 {
			// Skip zero values
			continue
		}
		out[wi] = v
		if wi != ri {
			// If we are writing to a different index,
			// it means a tile moved
			moved = true
		}
		wi++
	}

	// Zeros already in place for out[wi:]
	return out, moved
}

// mergeLine merges adjacent equal tiles in a line that's already been slid.
// Returns the merged+slid line and the total score gained.
// e.g.) [2, 2, 4, 0] -> [4, 4, 0, 0] (score gain = 2)
func mergeLine(line []int) ([]int, int) {
	n := len(line)
	scoreGain := 0

	for i := 0; i < n-1; i++ {
		if line[i] != 0 && line[i] == line[i+1] {
			line[i] *= 2  // Merge tiles
			line[i+1] = 0 // The second merged cell becomes zero
			scoreGain += line[i]
			i++
		}
	}

	// Final slide to remove zeros created by merging
	slid, _ := slideLine(line)
	return slid, scoreGain
}

// slideMergeLine does both in one shot: slide, merge, slide.
// Returns new line, whether anything moved/merged, and score gained.
func slideMergeLine(line []int) ([]int, bool, int) {
	slid, moved1 := slideLine(line)
	merged, gain := mergeLine(slid)
	_, moved2 := slideLine(merged)
	return merged, (moved1 || gain > 0 || moved2), gain
}
