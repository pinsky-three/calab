package board

import (
	"github.com/minskylab/calab"
)

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

	return &Board2D{
		dims:  []uint64{uint64(w), uint64(h)},
		Board: board,
		// totalSymbols: states,
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

func (s *Board2D) Fill(src Source2D, rule calab.Evolvable) *Board2D {
	totalStates := rule.Symbols()

	for i := int64(0); i < int64(s.dims[0]); i++ {
		for j := int64(0); j < int64(s.dims[1]); j++ {
			s.Board[i][j] = src(i, j, totalStates)
		}
	}

	return s
}
