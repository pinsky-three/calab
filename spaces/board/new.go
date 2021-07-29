package board

// New returns a new Board2D Space.
func New(w, h int) (*Board2D, error) {
	board := [][]uint64{}
	for i := int64(0); i < int64(w); i++ {
		r := []uint64{}
		for j := int64(0); j < int64(h); j++ {
			r = append(r, 0)
		}
		board = append(board, r)
	}

	// return board

	// board := initialState(int64(w), int64(h), src)
	dims := []uint64{uint64(len(board)), uint64(len(board[0]))}

	return &Board2D{
		Board: board,
		// totalSymbols: states,
		dims: dims,
	}, nil
}

// MustNew ...
func MustNew(w, h int) *Board2D {
	b, err := New(w, h)
	if err != nil {
		panic(err)
	}

	return b
}

func (s *Board2D) Fill(src Initial2DSource, states uint64) {
	// board := [][]uint64{}
	for i := int64(0); i < int64(s.dims[0]); i++ {
		r := []uint64{}
		for j := int64(0); j < int64(s.dims[1]); j++ {
			r = append(r, src(i, j, states))
		}

		s.Board = append(s.Board, r)
	}
}
