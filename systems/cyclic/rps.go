package cyclic

import (
	"math/rand"

	"github.com/minskylab/calab"
)

// RockPaperScissor return a rock-paper-scissor cellular automaton rules.
type RockPaperScissor struct {
	totalStates  uint64
	threshold    int
	stochastic   int
	neighborhood Neighborhood
	bounder      Bounder2D
}

func (prs *RockPaperScissor) counts(countsMap map[uint64]int, i, j int64, neighbors []uint64) {
	for s := range countsMap {
		countsMap[s] = 0
	}

	for _, n := range neighbors {
		countsMap[n]++
	}
}

func (prs *RockPaperScissor) Symbols() uint64 {
	return prs.totalStates
}

// Evolve perform the classic rule of the Rock Paper Scissor CA.
func (prs *RockPaperScissor) Evolve(space calab.Space) calab.Space {
	spaceState := space.Space()
	dims := space.Dims()

	board := [][]uint64{}

	countsMap := make(map[uint64]int)
	for s := uint64(0); s < prs.totalStates; s++ {
		countsMap[s] = 0
	}

	for j := uint64(0); j < uint64(dims[1]); j++ {
		board = append(board, spaceState[j*dims[1]:(j+1)*dims[1]])
	}

	newBoard := [][]uint64{}
	for _, b := range board {
		newBoard = append(newBoard, append([]uint64{}, b...))
	}

	for i := int64(0); i < int64(dims[0]); i++ {
		for j := int64(0); j < int64(dims[1]); j++ {
			prs.counts(countsMap, i, j, prs.neighborhood(&board, i, j, prs.bounder))

			nextState := (board[i][j] + 1) % prs.totalStates
			// if nextState > uint64(len(prs.countsMap))-1 {
			// 	nextState = 0
			// }

			th := prs.threshold
			if prs.stochastic > 0 {
				th += rand.Intn(prs.stochastic)
			}

			if countsMap[nextState] > th {
				newBoard[i][j] = nextState
			}
		}
	}

	nextSpace := []uint64{}
	for i := range newBoard {
		nextSpace = append(nextSpace, newBoard[i]...)
	}

	return space.Branch(nextSpace)
}
