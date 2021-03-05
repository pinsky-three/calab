package main

import (
	"bytes"
	"image/png"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/remote"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/spaces/board/renderers"
	"github.com/minskylab/calab/systems/lifelike"
)

func main() {
	// creating the space.
	width, height := 256, 256

	nh := board.MooreNeighborhood(2, false)
	bound := board.ToroidBounded()
	space := board.MustNew(width, height, nh, bound, board.RandomInit, board.UniformNoise)

	// creating the rule.
	rule := lifelike.MustNew(&lifelike.Rule{
		S: []int{3},
		B: []int{5, 8},
	})

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	// defining rendering
	c0, _ := colorful.Hex("#000000")
	c1, _ := colorful.Hex("#FFDF53")

	p := renderers.NewPalette(c0, c1, 2)
	imgRenderer := renderers.MustNewBoard(width, height, p)

	buff := bytes.NewBuffer([]byte{})
	dataBinary := make(chan []byte)

	srv := remote.NewBinaryRemote(3000, "/", dataBinary)

	vm := calab.NewVM(system, func(n uint64, s calab.Space) {
		img := imgRenderer.Render(n, s)

		buff.Reset()
		if err := png.Encode(buff, img); err != nil {
			panic(err)
		}

		dataBinary <- buff.Bytes()
	})

	go vm.Run(1000 * time.Second)

	srv.Run()
}
