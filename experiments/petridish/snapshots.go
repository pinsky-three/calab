package petridish

import (
	"image"

	"github.com/minskylab/calab"
)

// Snapshot perform a instant snapshot of your dynamical system.
func (pd *PetriDish) snapshot() {
	space := pd.Model.System.Space
	dims := space.Dims()

	w, h := dims[0], dims[1]

	for i := int64(0); i < int64(w); i++ {
		for j := int64(0); j < int64(h); j++ {
			pd.img.Set(int(i), int(j), pd.colorPalette[space.State(i, j)])
		}
	}
}

func (pd *PetriDish) renderImage(n uint64, s calab.Space) {
	// if !pd.Headless {
	// 	pd.snapshot()
	// }
	pd.ticks = n
	// pd.dataHole <- pd.buffer.Bytes()
}

// TakeSnapshot take a snapshot.
func (pd *PetriDish) TakeSnapshot() image.Image {
	pd.snapshot()
	return pd.img
}

// // SaveSnapshot save your snapshot in a file.
// func (pd *PetriDish) SaveSnapshot(name string, format SnapshotImageFormat) error {
// 	return pd.storage.SaveSnapshot(name, pd.img, format)
// }
