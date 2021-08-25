package cyclic

// Bounder2D represents a new bounder mapper
type Bounder2D func(w, h, xi, yi int64) (int64, int64)

// ToroidBounded is a toroind kind bounder
func ToroidBounded(w, h, xi, yi int64) (int64, int64) {
	if xi > int64(w-1) {
		xi = 0
	}

	if xi < 0 {
		xi = int64(w - 1)
	}

	if yi > int64(h-1) {
		yi = 0
	}

	if yi < 0 {
		yi = int64(h - 1)
	}

	return xi, yi
}

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

func VonNewmannNeighborhood(radius int, includeCenter bool) Neighborhood {
	return func(board *[][]uint64, x, y int64, bounder Bounder2D) []uint64 {
		nh := []uint64{}

		return nh
	}
}
