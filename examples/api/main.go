package main

import (
	"fmt"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/server"
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

	fmt.Printf("classicLifeLike id: %s\n", classicLifeLike.ID)
	fmt.Printf("rockPaperSicsors id: %s\n", rockPaperSicsors.ID)

	server.LaunchServer(experiment)
}
