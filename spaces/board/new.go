package board

// New returns a new Board2D Space.
func New(w, h int, neighborhood Neighborhood, bounder Bounder2D, initialState InitialStateGenerator, src Initial2DSource) (*Board2D, error) {
	board := initialState(int64(w), int64(h), src)
	dims := []uint64{uint64(len(board)), uint64(len(board[0]))}

	return &Board2D{
		Board:        board,
		neighborhood: neighborhood,
		bounder:      bounder,
		dims:         dims,
	}, nil
}

// MustNew ...
func MustNew(w, h int, neighborhood Neighborhood, bounder Bounder2D, initialState InitialStateGenerator, src Initial2DSource) *Board2D {
	b, err := New(w, h, neighborhood, bounder, initialState, src)
	if err != nil {
		panic(err)
	}

	return b
}
