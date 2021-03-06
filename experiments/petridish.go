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
	VM           *calab.VirtualMachine
	buffer       *bytes.Buffer
	img          *image.RGBA
	colorPalette calab.Palette
	dataHole     chan []byte
	ticks        uint64
	storage      ExperimentsStorage
}

// Snapshot perform a instant snapshot of your dynamical system.
func (pd *PetriDish) snapshot() {
	space := pd.VM.Model.Space
	dims := space.Dims()

	w, h := dims[0], dims[1]

	for i := int64(0); i < int64(w); i++ {
		for j := int64(0); j < int64(h); j++ {
			pd.img.Set(int(i), int(j), pd.colorPalette[space.State(i, j)])
		}
	}
}

func (pd *PetriDish) renderPNG(n uint64, s calab.Space) {
	pd.snapshot()

	pd.buffer.Reset()
	if err := png.Encode(pd.buffer, pd.img); err != nil {
		panic(err)
	}

	pd.ticks = n
	// pd.dataHole <- pd.buffer.Bytes()
}

// Run ...
func (pd *PetriDish) Run(duration time.Duration) {
	pd.VM.Run(duration)
}

// Ticks returns the current ticks in the model of your petri dish.
func (pd *PetriDish) Ticks() uint64 {
	return pd.ticks
}
