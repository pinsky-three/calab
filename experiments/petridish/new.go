package petridish

import (
	"bytes"
	"image"

	"github.com/minskylab/calab"
	uuid "github.com/satori/go.uuid"
)

const DefaultTPS = 1 << 20

func NewFromSystem(system *calab.DynamicalSystem) *PetriDish {
	vm := calab.NewVM(system)
	return NewFromVCM(vm)
}

// NewFromVCM returns a new petridish.
func NewFromVCM(vcm *calab.VirtualComputationalModel) *PetriDish {
	dims := vcm.Model.Space.Dims()

	pd := &PetriDish{}
	pd.ID = uuid.NewV4().String()

	pd.Model = vcm
	pd.colorPalette = calab.MonochromePalette(vcm.Model.Rule.Symbols())
	pd.buffer = bytes.NewBuffer([]byte{})
	pd.img = image.NewRGBA(image.Rect(0, 0, int(dims[0]), int(dims[1])))

	pd.Model.AddRenderer(pd.renderImage)

	pd.Model.Model.SetTPS(DefaultTPS)
	pd.Model.SetRPS(DefaultTPS)

	return pd
}
