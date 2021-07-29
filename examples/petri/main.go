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

	width, height := 2048, 2048

	rule := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))

	space := board.MustNew(width, height).Fill(board.UniformNoise, rule)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	// srv := remote.NewBinaryRemote(3000, "/", pd.binaryChannel)

	vm := calab.NewVM(system)

	pd := experiments.NewPetriDish(vm, calab.Palette{0: dead, 1: alive}, 50000)
	pd.Headless = true

	startTime := time.Now()

	done := pd.RunTicks(200)

	<-done

	f, err := os.OpenFile("snapshot.png", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	_ = png.Encode(f, pd.TakeSnapshot())

	fmt.Printf("experiment duration: %s | total ticks: %d\n", time.Since(startTime), pd.Ticks())

	f.Close()
}
