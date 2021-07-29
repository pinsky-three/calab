package experiments

import (
	"bytes"
	"image"

	"github.com/minskylab/calab"
	uuid "github.com/satori/go.uuid"
)

const DefaultTPS = 1 << 20

// NewPetriDish returns a new petridish.
func NewPetriDish(vcm *calab.VirtualComputationalModel, palette calab.Palette) *PetriDish {
	dims := vcm.Model.Space.Dims()

	pd := &PetriDish{}
	pd.ID = uuid.NewV4().String()

	pd.Model = vcm
	pd.colorPalette = palette
	pd.buffer = bytes.NewBuffer([]byte{})
	pd.img = image.NewRGBA(image.Rect(0, 0, int(dims[0]), int(dims[1])))

	pd.Model.AddRenderer(pd.renderImage)

	pd.Model.Model.SetTPS(DefaultTPS)
	pd.Model.SetRPS(DefaultTPS)

	return pd
}
