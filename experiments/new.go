package experiments

import (
	"bytes"
	"image"

	"github.com/minskylab/calab"
)

// NewPetriDish returns a new petridish.
func NewPetriDish(vm *calab.VirtualMachine, palette calab.Palette, tps int) *PetriDish {
	dims := vm.Model.Space.Dims()

	pd := &PetriDish{}

	pd.VM = vm
	pd.colorPalette = palette
	pd.buffer = bytes.NewBuffer([]byte{})
	pd.img = image.NewRGBA(image.Rect(0, 0, int(dims[0]), int(dims[1])))

	pd.VM.AddRenderer(pd.renderImage)

	pd.VM.Model.SetTPS(tps)
	pd.VM.SetRPS(tps)

	return pd
}
