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

// PetriDish is a petri dish.
type PetriDish struct {
	Width, Height int
	palette       calab.Palette
	imgRenderer   *renderers.BoardImageRenderer
	buffer        *bytes.Buffer
	binaryChannel chan []byte
}

func main() {
	var c0, _ = colorful.Hex("#1e2031")
	var c1, _ = colorful.Hex("#fbe3a1")

	width, height := 168, 56

	palette := calab.Palette{0: c1, 1: c0, 2: c1, 3: c0}

	pd := PetriDish{
		Width:  width,
		Height: height,

		palette:       palette,
		imgRenderer:   renderers.MustNewBoard(width, height, palette),
		buffer:        bytes.NewBuffer([]byte{}),
		binaryChannel: make(chan []byte),
	}

	// creating the space.
	nh := board.MooreNeighborhood(1, false)
	bound := board.ToroidBounded()
	space := board.MustNew(width, height, nh, bound, board.RandomInit, board.UniformNoise(len(palette)))

	// creating the rule.
	// rule := lifelike.MustNew(lifelike.DayAndNight)
	rule := cyclic.MustNewRockPaperScissor(len(palette), 1, 4)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	srv := remote.NewBinaryRemote(3000, "/", pd.binaryChannel)

	vm := calab.NewVM(system, pd.renderImage)

	go vm.Run(1000 * time.Second)

	srv.Run()
}

func (pd *PetriDish) renderImage(n uint64, s calab.Space) {
	img := pd.imgRenderer.Render(n, s)

	pd.buffer.Reset()
	if err := jpeg.Encode(pd.buffer, img, &jpeg.Options{
		Quality: 100,
	}); err != nil {
		panic(err)
	}

	pd.binaryChannel <- pd.buffer.Bytes()
}
