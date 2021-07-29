package board

import (
	"github.com/minskylab/calab"
)

// Board2D represents a 2D Board Space
type Board2D struct {
	Board [][]uint64
	dims  []uint64
}

// State implements DS Space interface.
func (s *Board2D) State(i ...int64) uint64 {
	return s.Board[i[0]][i[1]]
}

func (s *Board2D) Space() []uint64 {
	space := []uint64{}

	for i := range s.Board {
		space = append(space, s.Board[i]...)
	}

	return space
}

// Dims return the dimension of 2D board.
func (s *Board2D) Dims() []uint64 {
	return s.dims
}

func (s *Board2D) Branch(space []uint64) calab.Space {
	board := [][]uint64{}
	for j := uint64(0); j < uint64(s.dims[1]); j++ {
		board = append(board, space[j*s.dims[1]:(j+1)*s.dims[1]])
	}

	s.Board = board

	return s
}
