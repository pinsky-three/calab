package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
)

func main() {
	dead, _ := colorful.Hex("#0e1717")
	alive, _ := colorful.Hex("#fbe722")

	width, height := 256, 256

	rule := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))

	space := board.MustNew(width, height).Fill(board.UniformNoise, rule)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	// srv := remote.NewBinaryRemote(3000, "/", pd.binaryChannel)

	vm := calab.NewVM(system)

	pd := experiments.NewPetriDish(vm, calab.Palette{0: dead, 1: alive}, 300)
	pd.Headless = true

	startTime := time.Now()

	done := pd.RunTicks(100)

	<-done

	expDuration := time.Since(startTime)

	f, err := os.OpenFile("snapshot.png", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	_ = png.Encode(f, pd.TakeSnapshot())

	fmt.Printf("experiment duration: %s | total ticks: %d | tps: %.2f\n", expDuration, pd.Ticks(), float64(pd.Ticks())/expDuration.Seconds())

	f.Close()
}
