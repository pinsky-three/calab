package board

import "math/rand"

// New returns a new Board2D Space.
func New(w, h int, neighborhood Neighborhood, bounder Bounder) (*Board2D, error) {
	board := [][]byte{}
	for i := 0; i < w; i++ {
		r := []byte{}
		for j := 0; j < h; j++ {
			r = append(r, byte(rand.Intn(2)))
		}
		board = append(board, r)
	}

	return &Board2D{
		Board:        board,
		neighborhood: neighborhood,
		bounder:      bounder,
	}, nil
}

// MustNew ...
func MustNew(w, h int, neighborhood Neighborhood, bounder Bounder) *Board2D {
	b, err := New(w, h, neighborhood, bounder)
	if err != nil {
		panic(err)
	}

	return b
}
