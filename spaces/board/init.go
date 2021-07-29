package board

import "math/rand"

// Source2D represents a simple 2d initial state supplier.
type Source2D func(x int64, y int64, states uint64) uint64

// UniformNoise is a basic uniform noise.]
func UniformNoise(x, y int64, states uint64) uint64 {
	return uint64(rand.Intn(int(states)))
}
