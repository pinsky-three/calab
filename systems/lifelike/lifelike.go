package lifelike

import (
	"github.com/minskylab/calab"
)

// Rule wraps a basic life like cellular automaton rule.
type Rule struct {
	S []int
	B []int
}

// LifeLike represents a simple life like cellular automaton rule.
type LifeLike struct {
	// board *board.Board2D
	neighborhood Neighborhood
	bounder      Bounder2D
	rule         *Rule
}

func (life *LifeLike) Symbols() uint64 {
	return 2
}

// Evolve implements a evolvable system for calab framework.
func (life *LifeLike) Evolve(space calab.Space) calab.Space {
	// board := space.(*board.Board2D)
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
			// rule evaluation
			neighborsSum := int(life.sum(life.neighborhood(&board, i, j, life.bounder)))
			life.calculateNextState(i, j, neighborsSum, &newBoard)
		}
	}

	nextSpace := make([]uint64, dims[0]*dims[1])
	for i := range newBoard {
		nextSpace = append(nextSpace, newBoard[i]...)
	}

	return space.Branch(nextSpace)
}

func (life *LifeLike) calculateNextState(i, j int64, neighborsSum int, board *[][]uint64) {
	for _, b := range life.rule.B {
		if b == neighborsSum {
			(*board)[i][j] = 1
			return
		}
	}

	for _, s := range life.rule.S {
		if s == neighborsSum {
			return
		}
	}

	(*board)[i][j] = 0
}

func (life *LifeLike) sum(states []uint64) uint64 {
	sum := uint64(0)
	for _, s := range states {
		sum += s
	}

	return sum
}
