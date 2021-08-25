package board

import "math/rand"

// Source2D represents a simple 2d initial state supplier.
type Source2D func(x int64, y int64, states uint64) (bool, uint64)

// UniformNoise is a basic uniform noise.]
func UniformNoise(x, y int64, states uint64) (bool, uint64) {
	return true, uint64(rand.Intn(int(states)))
}

func FullState(state uint64) Source2D {
	return func(x, y int64, states uint64) (bool, uint64) {
		return true, state
	}
}

func SpecificPositions(statesMapping map[uint64][][]int) Source2D {
	return func(x, y int64, states uint64) (bool, uint64) {
		for state, position := range statesMapping {
			for _, pos := range position {
				if int64(pos[0]) == x && int64(pos[1]) == y {
					return true, state
				}
			}
		}
		return false, 0
	}
}
