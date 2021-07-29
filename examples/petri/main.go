package main

import (
	"fmt"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
	"github.com/minskylab/calab/utils"
)

func main() {
	dead, _ := colorful.Hex("#0e1717")
	alive, _ := colorful.Hex("#fbe722")

	width, height := 512, 512

	rule := lifelike.MustNew(lifelike.HighLifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))

	space := board.MustNew(width, height).Fill(board.UniformNoise, rule)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	vm := calab.NewVM(system)

	pd := experiments.NewPetriDish(vm, calab.Palette{0: dead, 1: alive})
	pd.Headless = false

	startTime := time.Now()

	<-pd.RunTicks(500)

	expDuration := time.Since(startTime)
	totalTicks := pd.Ticks()
	tps := float64(totalTicks) / expDuration.Seconds()

	fmt.Printf("experiment duration: %s | total ticks: %d | tps: %.2f\n", expDuration, totalTicks, tps)

	if err := utils.SaveSnapshotAsPNG(pd, pd.ID+".png"); err != nil {
		panic(err)
	}
}
