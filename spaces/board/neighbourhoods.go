package board

// Neighborhood is a function for generic neighborhoods.
type Neighborhood func(board *[][]byte, x, y int64, bounder Bounder) []uint64

// // MooreNeighborhood is a versatil (with radius and auto inclusion) Moore Neighborhood.
// func MooreNeighborhood(board *[][]byte, w, h, x, y int64, r int, excludeItself bool, bounder BoardBounder) []uint64 {

// }

// MooreNeighborhood is a versatil (with radius and auto inclusion) Moore Neighborhood.
func MooreNeighborhood(radius int, includeCenter bool) Neighborhood {
	return func(board *[][]byte, x, y int64, bounder Bounder) []uint64 {
		nh := []uint64{}

		for dx := int64(-radius); dx < int64(radius+1); dx++ {
			for dy := int64(-radius); dy < int64(radius+1); dy++ {
				xi := x + dx
				yi := y + dy

				if xi == x && yi == y && !includeCenter {
					continue
				}

				xi, yi = bounder(int64(len(*board)), int64(len((*board)[0])), xi, yi)

				// s.reusableNeighbors = append(s.reusableNeighbors, uint64(s.Board[xi][yi]))
				nh = append(nh, uint64((*board)[xi][yi]))
			}
		}

		return nh
	}
}
