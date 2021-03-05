package cyclic

import (
	"math/rand"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/spaces/board"
)

// RockPaperScissor return a rock-paper-scissor cellular automaton rules.
type RockPaperScissor struct {
	countsMap  map[uint64]int
	threshold  int
	stochastic int
}

func (prs *RockPaperScissor) counts(i, j int64, neighbors []uint64) {
	for s := range prs.countsMap {
		prs.countsMap[s] = 0
	}

	for _, n := range neighbors {
		prs.countsMap[n]++
	}
}

// Evolve perform the classic rule of the Rock Paper Scissor CA.
func (prs *RockPaperScissor) Evolve(space calab.Space) {
	board := space.(*board.Board2D)
	dims := space.Dims()

	newBoard := [][]uint64{}
	for _, b := range board.Board {
		newBoard = append(newBoard, append([]uint64{}, b...))
	}

	for i := int64(0); i < int64(dims[0]); i++ {
		for j := int64(0); j < int64(dims[1]); j++ {
			prs.counts(i, j, board.Neighbors(i, j))

			nextState := board.Board[i][j] + 1
			if nextState > uint64(len(prs.countsMap))-1 {
				nextState = 0
			}

			th := prs.threshold
			if prs.stochastic > 0 {
				th += rand.Intn(prs.stochastic)
			}

			if prs.countsMap[nextState] > th {
				newBoard[i][j] = nextState
			}
		}
	}

	board.Board = newBoard
}
