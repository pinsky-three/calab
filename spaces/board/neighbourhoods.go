package board

// Neighborhood is a function for generic neighborhoods.
type Neighborhood func(board *[][]uint64, x, y int64, bounder Bounder2D) []uint64

// MooreNeighborhood is a versatil (with radius and auto inclusion) Moore Neighborhood.
func MooreNeighborhood(radius int, includeCenter bool) Neighborhood {
	return func(board *[][]uint64, x, y int64, bounder Bounder2D) []uint64 {
		nh := []uint64{}
		w, h := int64(len(*board)), int64(len((*board)[0]))

		for dx := int64(-radius); dx < int64(radius+1); dx++ {
			for dy := int64(-radius); dy < int64(radius+1); dy++ {
				xi := x + dx
				yi := y + dy

				if xi == x && yi == y && !includeCenter {
					continue
				}

				xi, yi = bounder(w, h, xi, yi)

				// s.reusableNeighbors = append(s.reusableNeighbors, uint64(s.Board[xi][yi]))
				nh = append(nh, uint64((*board)[xi][yi]))
			}
		}

		return nh
	}
}
