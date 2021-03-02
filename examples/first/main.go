package main

import (
	"fmt"
	"time"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
)

func main() {
	// creating the space.
	width, height := 512, 512
	nh := board.MooreNeighborhood(1, false)
	space := board.MustNew(width, height, nh, board.ToroidBounded)

	// creating the rule.
	rule := lifelike.MustNew(lifelike.GameOfLifeRule)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	// ticks := make(chan uint64)

	// system.RunSimulation(100, ticks)

	fmt.Println(system.ID)

	time.Sleep(10 * time.Second)
}
