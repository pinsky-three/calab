package main

import (
	"fmt"
	"time"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/experiments/utils"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/cyclic"
	"github.com/minskylab/calab/systems/lifelike"
)

func basicLifeLike(w, h int, lifeRule *lifelike.Rule) *petridish.PetriDish {
	dynamic := lifelike.MustNew(lifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))
	space := board.MustNew(w, h).Fill(board.UniformNoise, dynamic)
	return petridish.NewFromSystem(calab.BulkDynamicalSystem(space, dynamic))
}

func fastCyclicAutomata(w, h int, radius, states, threshold, stochastic int) *petridish.PetriDish {
	nh := cyclic.MooreNeighborhood(radius, false)
	dynamic := cyclic.MustNewRockPaperScissor(cyclic.ToroidBounded, nh, states, threshold, stochastic)

	space := board.MustNew(w, h).Fill(board.UniformNoise, dynamic)
	return petridish.NewFromSystem(calab.BulkDynamicalSystem(space, dynamic))
}

func main() {
	classicLifeLike := basicLifeLike(256, 256, lifelike.GameOfLifeRule)
	rockPaperSicsors := fastCyclicAutomata(256, 256, 2, 7, 2, 1)

	experiment := experiments.New()

	experiment.AddPetriDish(classicLifeLike)
	experiment.AddPetriDish(rockPaperSicsors)

	experiment.Run(classicLifeLike.ID, experiments.WithTicks(1000))
	experiment.Run(rockPaperSicsors.ID, experiments.WithTicks(1000))

	time.Sleep(20 * time.Second)

	fmt.Printf("classicLifeLike ends with %d ticks\n", classicLifeLike.Ticks())
	fmt.Printf("rockPaperSicsors ends with %d ticks\n", rockPaperSicsors.Ticks())

	utils.SaveSnapshotAsPNG(classicLifeLike, "classicLifeLike.png")
	utils.SaveSnapshotAsPNG(rockPaperSicsors, "rockPaperSicsors.png")
}
