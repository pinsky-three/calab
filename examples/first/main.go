package main

import (
	"bytes"
	"image/jpeg"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/remote"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/spaces/board/renderers"

	// "github.com/minskylab/calab/systems/lifelike"
	"github.com/minskylab/calab/systems/cyclic"
)

const width, height = 168, 56
const states = 4

var c0, _ = colorful.Hex("#1e2031")
var c1, _ = colorful.Hex("#faad65")

// var p = renderers.NewPalette(c0, c1, states)
var p = renderers.Palette{
	0: c1,
	1: c0,
	2: c1,
	3: c0,
}

var imgRenderer = renderers.MustNewBoard(width, height, p)

var buff = bytes.NewBuffer([]byte{})

var dataBinary = make(chan []byte)

func main() {
	// creating the space.

	nh := board.MooreNeighborhood(1, false)
	bound := board.ToroidBounded()
	space := board.MustNew(width, height, nh, bound, board.RandomInit, board.UniformNoise(states))

	// creating the rule.
	// rule := lifelike.MustNew(lifelike.DayAndNight)
	rule := cyclic.MustNewRockPaperScissor(states, 1, 4)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule, 60)

	srv := remote.NewBinaryRemote(3000, "/", dataBinary)

	vm := calab.NewVM(system, 60, renderImage)

	go vm.Run(1000 * time.Second)

	srv.Run()
}

func renderImage(n uint64, s calab.Space) {
	img := imgRenderer.Render(n, s)

	buff.Reset()
	if err := jpeg.Encode(buff, img, &jpeg.Options{
		Quality: 100,
	}); err != nil {
		panic(err)
	}

	dataBinary <- buff.Bytes()
}
