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
	"github.com/minskylab/calab/systems/cyclic"
)

func main() {

	var c0, _ = colorful.Hex("#0e1717")
	var c1, _ = colorful.Hex("#fbe722")

	width, height := 256, 256

	// palette := calab.Palette{0: c1, 1: c0, 2: c1}
	palette := calab.NewPalette(c0, c1, 3)

	// creating the space.
	nh := board.MooreNeighborhood(1, false)
	bound := board.ToroidBounded()
	space := board.MustNew(width, height, nh, bound, board.RandomInit, board.UniformNoise(len(palette)))

	// creating the rule.
	// rule := lifelike.MustNew(lifelike.DayAndNight)
	rule := cyclic.MustNewRockPaperScissor(len(palette), 2, 1)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	// srv := remote.NewBinaryRemote(3000, "/", pd.binaryChannel)

	vm := calab.NewVM(system)

	pd := experiments.NewPetriDish(vm, palette, 50000)
	pd.Headless = true

	startTime := time.Now()

	done := pd.RunTicks(2000)

	<-done

	fmt.Printf("experiment duration: %s | total ticks: %d\n", time.Since(startTime), pd.Ticks())

	f, err := os.OpenFile("snapshot.png", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	_ = png.Encode(f, pd.TakeSnapshot())

	f.Close()
}
