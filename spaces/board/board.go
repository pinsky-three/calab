package board

// Board2D represents a 2D Board Space
type Board2D struct {
	Board        [][]byte
	neighborhood Neighborhood
	bounder      Bounder

	// reusableNeighbors []uint64
}

// State implements DS Space interface.
func (s *Board2D) State(i ...int64) uint64 {
	return uint64(s.Board[i[0]][i[1]])
}

// Neighbors implements DS Space interface.
func (s *Board2D) Neighbors(i ...int64) []uint64 {
	x, y := i[0], i[1]

	// s.reusableNeighbors = []uint64{}

	return s.neighborhood(&s.Board, x, y, s.bounder)
	// s.mooreNeighborhood(x, y, 1, s.toroidBounded)

	// return s.reusableNeighbors
}

// Dims return the dimension of 2D board.
func (s *Board2D) Dims() []uint64 {
	return []uint64{uint64(len(s.Board)), uint64(len(s.Board[0]))}
}
