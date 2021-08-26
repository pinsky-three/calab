package main

import (
	"fmt"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/cyclic"
)

func cellularAutomaton(w, h int, radius, states, threshold, stochastic int) *petridish.PetriDish {
	nh := cyclic.MooreNeighborhood(radius, false)
	dynamic := cyclic.MustNewRockPaperScissor(cyclic.ToroidBounded, nh, states, threshold, stochastic)

	space := board.MustNew(w, h).Fill(board.UniformNoise, dynamic)
	return petridish.NewFromSystem(calab.BulkDynamicalSystem(space, dynamic), petridish.WithTPSMonitor)
}

func generateTimelapse(experiment *experiments.Experiment, r, s, t int) {
	ca := cellularAutomaton(512, 512, r, s, t, 1)
	experiment.AddPetriDish(ca)

	fmt.Printf("added CA: %s\n", ca.ID)

	timelapseOptions := &experiments.TimeLapseOptions{
		Debug:          true,
		DeleteAfter:    true,
		OutputFilename: fmt.Sprintf("results/cyclic-%d-%d-%d.mp4", r, s, t),
	}

	done := ca.RunTicks(100)
	if err := experiment.Timelapse(ca.ID, done, timelapseOptions); err != nil {
		panic(err)
	}
}

func main() {
	experiment := experiments.New()

	radius := []int{1, 2}
	stages := []int{3, 4, 6}
	thresholds := []int{2, 3, 4}

	for _, s := range stages {
		for _, r := range radius {
			for _, t := range thresholds {
				go generateTimelapse(experiment, r, s, t)
			}
		}
	}

	// generateTimelapse(experiment, 10, 2, 2)
}
