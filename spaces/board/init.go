package board

import "math/rand"

// Initial2DSource represents a simple 2d initial state supplier.
type Initial2DSource func(x int64, y int64, states uint64) uint64

// UniformNoise is a basic uniform noise.]
func UniformNoise(x, y int64, states uint64) uint64 {
	return uint64(rand.Intn(int(states)))
}
