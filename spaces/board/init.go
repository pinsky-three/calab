package board

import "math/rand"

// Initial2DSource represents a simple 2d initial state supplier.
type Initial2DSource func(x int64, y int64) uint64

// InitialStateGenerator  is a initial state generator.
type InitialStateGenerator func(w, h int64, source Initial2DSource) [][]uint64

// RandomInit is a random initial source.
func RandomInit(w, h int64, source Initial2DSource) [][]uint64 {
	board := [][]uint64{}
	for i := int64(0); i < w; i++ {
		r := []uint64{}
		for j := int64(0); j < h; j++ {
			r = append(r, source(i, j))
		}

		board = append(board, r)
	}

	return board
}

// UniformNoise is a basic uniform noise.]
func UniformNoise(states int) Initial2DSource {
	return func(x, y int64) uint64 {
		return uint64(rand.Intn(states))
	}
}
