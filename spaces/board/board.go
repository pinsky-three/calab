package board

// Board2D represents a 2D Board Space
type Board2D struct {
	Board        [][]uint64
	neighborhood Neighborhood
	bounder      Bounder2D
	dims         []uint64
}

// State implements DS Space interface.
func (s *Board2D) State(i ...int64) uint64 {
	return s.Board[i[0]][i[1]]
}

// Neighbors implements DS Space interface.
func (s *Board2D) Neighbors(i ...int64) []uint64 {
	x, y := i[0], i[1]

	return s.neighborhood(&s.Board, x, y, s.bounder)
}

// Dims return the dimension of 2D board.
func (s *Board2D) Dims() []uint64 {
	return s.dims
}
