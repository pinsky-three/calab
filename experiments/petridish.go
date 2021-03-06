package experiments

import (
	"bytes"
	"image"
	"image/png"
	"time"

	"github.com/minskylab/calab"
)

// PetriDish represents a fully instrumented system.
type PetriDish struct {
	vm           *calab.VirtualMachine
	buffer       *bytes.Buffer
	img          *image.RGBA
	colorPalette calab.Palette
	dataHole     chan []byte
}

// Snapshot perform a instant snapshot of your dynamical system.
func (pd *PetriDish) Snapshot() image.Image {
	space := pd.vm.Model.Space
	dims := space.Dims()

	w, h := dims[0], dims[1]

	for i := int64(0); i < int64(w); i++ {
		for j := int64(0); j < int64(h); j++ {
			pd.img.Set(int(i), int(j), pd.colorPalette[space.State(i, j)])
		}
	}

	return pd.img
}

func (pd *PetriDish) renderPNG(n uint64, s calab.Space) {
	pd.Snapshot()

	pd.buffer.Reset()
	if err := png.Encode(pd.buffer, pd.img); err != nil {
		panic(err)
	}

	// pd.dataHole <- pd.buffer.Bytes()
}

// Run ...
func (pd *PetriDish) Run(duration time.Duration) {
	go pd.vm.Run(duration)
}
