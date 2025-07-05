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

// mergeLine merges adjacent tiles with the same value.
// But, it does not slide the tiles.
// e.g.) [2, 2, 4, 0] -> [4, 0, 4, 0] (scoreGain: 4)
func mergeLine(line []int) ([]int, int, bool) {
	n := len(line)
	scoreGain := 0
	merged := false

	for i := 0; i < n-1; i++ {
		if line[i] != 0 && line[i] == line[i+1] {
			// merge
			line[i] *= 2
			line[i+1] = 0
			scoreGain += line[i]
			merged = true

			i++
		}
	}

	// Don't slide here, the caller will do it
	return line, scoreGain, merged
}

// slideMergeLine combines sliding and merging in one step.
// It first slides the line, then merges adjacent tiles,
// and finally slides again to compact the line.
// Returns the final line, a boolean indicating if any tile moved,
func slideMergeLine(line []int) ([]int, bool, int) {
	// 1. initial slide
	slid, moved1 := slideLine(line)

	// 2. merge (Without the final slide inside it)
	merged, scoreGain, didMerge := mergeLine(slid)

	// 3. final slide to compact the line
	final, _ := slideLine(merged)

	// The move is successful if the initial slide did something or
	// if a merge happened
	return final, (moved1 || didMerge), scoreGain
}
