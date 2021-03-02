package cyclic

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/minskylab/calab"
)

// PaperRockSissor is a simple cyclic automaton.
type PaperRockSissor struct {
	board        [][]byte
	countsMap    map[byte]int
	threshold    int
	colormap     map[byte]color.Color
	frame        *image.RGBA
	images       chan image.Image
	stocasticity int
}

// State ...
func (prs *PaperRockSissor) State(i ...int64) uint64 {
	return uint64(prs.board[i[0]][i[1]])
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// Neighbours ...
func (prs *PaperRockSissor) Neighbours(i ...int64) []uint64 {
	x, y := i[0], i[1]

	ns := []uint64{}

	for dx := int64(-3); dx < 4; dx++ {
		for dy := int64(-3); dy < 4; dy++ {
			if abs(dx) < 2 || abs(dy) < 2 {
				continue
			}

			xi := x + dx
			yi := y + dy

			if xi == x && yi == y {
				continue
			}

			if xi > int64(len(prs.board)-1) || xi < 0 {
				continue
			}

			if yi > int64(len(prs.board[0])-1) || yi < 0 {
				continue
			}

			ns = append(ns, uint64(prs.board[xi][yi]))
		}
	}

	return ns
}

func (prs *PaperRockSissor) counts(i, j int64) {
	for s := range prs.countsMap {
		prs.countsMap[s] = 0
	}

	for _, n := range prs.Neighbours(i, j) {
		prs.countsMap[byte(n)]++
	}

}

// Evolve ...
func (prs *PaperRockSissor) Evolve(space calab.Space) calab.Space {
	newBoard := make([][]byte, len(prs.board))
	for i := range newBoard {
		newBoard[i] = make([]byte, len(prs.board[0]))
	}

	for i := int64(0); i < int64(len(prs.board)); i++ {
		for j := int64(0); j < int64(len(prs.board[0])); j++ {
			prs.counts(i, j)
			// prs.countsMap[]

			nextState := prs.board[i][j] + 1
			if nextState > byte(len(prs.colormap))-1 {
				nextState = 0
			}

			th := prs.threshold
			if prs.stocasticity > 0 {
				th += rand.Intn(prs.stocasticity)
			}

			if prs.countsMap[nextState] > th {
				newBoard[i][j] = nextState
			} else {
				newBoard[i][j] = prs.board[i][j]
			}
		}
	}

	prs.board = newBoard

	return prs
}

// Render ...
func (prs *PaperRockSissor) Render(n uint64, s calab.Space) {
	for i := int64(0); i < int64(len(prs.board)); i++ {
		for j := int64(0); j < int64(len(prs.board[0])); j++ {
			prs.frame.Set(int(i), int(j), prs.colormap[prs.board[i][j]])
		}
	}

	prs.images <- prs.frame
}
