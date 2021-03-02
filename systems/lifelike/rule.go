package lifelike

import (
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/spaces/board"
)

// Rule wraps a basic life like cellular automaton rule.
type Rule struct {
	S []int
	B []int
}

// LifeLike represents a simple life like cellular automaton rule.
type LifeLike struct {
	// board *board.Board2D
	rule *Rule
}

// Evolve implements a evolvable system for calab framework.
func (life *LifeLike) Evolve(space calab.Space) {
	board := space.(*board.Board2D)
	dims := board.Dims()

	newBoard := [][]byte{}
	for _, b := range board.Board {
		newBoard = append(newBoard, b)
	}

	for i := int64(0); i < int64(dims[0]); i++ {
		for j := int64(0); j < int64(dims[1]); j++ {
			// rule evaluation
			neighborsSum := int(life.sum(board.Neighbors(i, j)))
			life.calculateNextState(i, j, neighborsSum, &newBoard)
		}
	}

	board.Board = newBoard
}

func (life *LifeLike) calculateNextState(i, j int64, neighborsSum int, board *[][]byte) {
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
