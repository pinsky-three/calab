package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/spaces/board/renderers"
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

	p := renderers.NewPallette(colorful.WarmColor(), colorful.HappyColor(), 2)

	imgRenderer, _ := renderers.NewBoardImage(width, height, p)

	ticks := make(chan uint64)
	done := make(chan struct{})

	f, err := os.OpenFile("frame.png", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	go func() {
		system.Observe(ticks, func(n uint64, s calab.Space) {
			img := imgRenderer.Render(n, s)
			fmt.Println("render", n)
			if err = png.Encode(f, img); err != nil {
				panic(err)
			}
		})
	}()

	system.RunInfiniteSimulation(ticks, done)
	<-done
}
