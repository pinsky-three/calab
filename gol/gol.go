package gol

import (
	"math/rand"

	"github.com/minskylab/calab"
)

// CA implements a Game Of Life CA for DynamicalSystem model.
type CA struct {
	board    [][]byte
	cellSize int
}

// State implements DS Space interface.
func (ca *CA) State(i ...int64) uint64 {
	return uint64(ca.board[i[0]][i[1]])
}

// Neighbors implements DS Space interface.
func (ca *CA) Neighbors(i ...int64) []uint64 {
	x, y := i[0], i[1]

	ns := []uint64{}

	for dx := int64(-1); dx < 2; dx++ {
		for dy := int64(-1); dy < 2; dy++ {
			xi := x + dx
			yi := y + dy

			if xi == x && yi == y {
				continue
			}

			if xi > int64(len(ca.board)-1) || xi < 0 {
				continue
			}

			if yi > int64(len(ca.board[0])-1) || yi < 0 {
				continue
			}

			ns = append(ns, uint64(ca.board[xi][yi]))
		}
	}

	return ns
}

// Evolve implements a Evolvable interface.
func (ca *CA) Evolve(space calab.Space) calab.Space {
	newBoard := make([][]byte, len(ca.board))
	for i := range newBoard {
		newBoard[i] = make([]byte, len(ca.board[0]))
	}

	for i := int64(0); i < int64(len(ca.board)); i++ {
		for j := int64(0); j < int64(len(ca.board[i])); j++ {
			total := 0
			for _, t := range space.Neighbors(i, j) {
				total += int(t)
			}

			if total == 2 {
				newBoard[i][j] = ca.board[i][j]
				continue
			}

			if total == 3 || total == 7 {
				newBoard[i][j] = 1
				continue
			}

			newBoard[i][j] = 0
		}
	}

	ca.board = newBoard

	return ca
}

// new creates a new GoL system.
func new(w, h int, randomSeed int64) *CA {
	rand.Seed(randomSeed)

	board := make([][]byte, h)
	for i := range board {
		board[i] = make([]byte, w)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = byte(rand.Intn(2))
		}
	}

	return &CA{
		board:    board,
		cellSize: 4,
	}
}

// NewGoLDynamicalSystem returns a new GoL DS.
func NewGoLDynamicalSystem(w, h int, randomSeed int64) *calab.DynamicalSystem {
	gol := new(w, h, randomSeed)
	return calab.BulkDynamicalSystem(gol, gol)
}
