package main

import (
	"fmt"
	"time"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/experiments/utils"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
)

func main() {
	width, height := 512, 512

	rule := lifelike.MustNew(lifelike.HighLifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))
	space := board.MustNew(width, height).Fill(board.UniformNoise, rule)
	system := calab.BulkDynamicalSystem(space, rule)

	pd := petridish.NewFromSystem(system)

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
