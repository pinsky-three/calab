package voronoi

import (
	"github.com/minskylab/calab"
)

func (v *Voronoi) bounder(w, h, xi, yi int64) (int64, int64) {
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

func (v *Voronoi) neighborhood(board *[][]uint64, x, y int64) []uint64 {
	nh := []uint64{}
	radius := 1
	w, h := int64(len(*board)), int64(len((*board)[0]))

	for dx := int64(-radius); dx < int64(radius+1); dx++ {
		for dy := int64(-radius); dy < int64(radius+1); dy++ {
			xi := x + dx
			yi := y + dy

			if xi == x && yi == y {
				continue
			}

			xi, yi = v.bounder(w, h, xi, yi)

			// s.reusableNeighbors = append(s.reusableNeighbors, uint64(s.Board[xi][yi]))

			nh = append(nh, (*board)[xi][yi])
		}
	}

	return nh
}

func (v *Voronoi) counts(states []uint64) map[uint64]int {
	counts := map[uint64]int{}

	for _, s := range states {
		counts[s] = counts[s] + 1
	}

	return counts
}

func (v *Voronoi) Evolve(space calab.Space) calab.Space {
	spaceState := space.Space()
	dims := space.Dims()

	board := [][]uint64{}

	for j := uint64(0); j < uint64(dims[1]); j++ {
		board = append(board, spaceState[j*dims[1]:(j+1)*dims[1]])
	}

	newBoard := [][]uint64{}
	for _, b := range board {
		newBoard = append(newBoard, append([]uint64{}, b...))
	}

	for i := int64(0); i < int64(dims[0]); i++ {
		for j := int64(0); j < int64(dims[1]); j++ {
			encodeNeighborhood := v.neighborhood(&board, i, j)

			// fmt.Println(encodeNeighborhood)

			neighborhoodCounts := v.counts(encodeNeighborhood)

			// if  {
			if _, ok := neighborhoodCounts[0]; len(neighborhoodCounts) == 1 && ok {
				continue
				// newBoard[i][j] = 0
			}
			// }

			// if len(neighborhoodCounts) > 2 {
			// 	continue
			// }

			// if isUniqueSymbolComposition(encodeNeighborhood) {
			// 	continue
			// }

			// for s := 1; s < v.totalStates+1; s++ {
			for s := uint64(0); s < uint64(v.totalPoints); s++ {
				if s == 0 {
					continue
				}

				total, have := neighborhoodCounts[s]

				if !have {
					continue
				}

				// if neighborhoodSum%s == 0 {
				// 	continue
				// }

				if total == 3 {
					newBoard[i][j] = s
					continue
				}

				if total > 4 {
					newBoard[i][j] = 0
					continue
				}
			}

			// newBoard[i][j] = 0
			// }

			// if neighborhoodSum == 3 {
			// 	newBoard[i][j] = 1
			// }

			// if neighborhoodSum > 4 {
			// 	newBoard[i][j] = 0
			// }

		}
	}

	nextSpace := []uint64{}
	for i := range newBoard {
		nextSpace = append(nextSpace, newBoard[i]...)
	}

	return space.Branch(nextSpace)
}

func (v *Voronoi) Symbols() uint64 {
	return uint64(v.totalPoints + 1)
}
