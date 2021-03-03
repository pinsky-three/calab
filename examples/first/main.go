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
	bound := board.ToroidBounded()
	space := board.MustNew(width, height, nh, bound, board.RandomInit, board.UniformNoise)

	// creating the rule.
	rule := lifelike.MustNew(lifelike.GameOfLifeRule)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	ticks := make(chan uint64)

	go func() {
		lastTime := time.Now()
		system.Observe(ticks, func(n uint64, s calab.Space) {
			fmt.Println(time.Since(lastTime))
			lastTime = time.Now()
		})

	}()

	fmt.Println(system.ID)

	system.RunSimulation(100, ticks)

	time.Sleep(10 * time.Second)
}
