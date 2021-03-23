package experiments

import (
	"bytes"
	"image"

	"github.com/minskylab/calab"
	uuid "github.com/satori/go.uuid"
)

// NewPetriDish returns a new petridish.
func NewPetriDish(vcm *calab.VirtualComputationalModel, palette calab.Palette, tps int) *PetriDish {
	dims := vcm.Model.Space.Dims()

	pd := &PetriDish{}
	pd.ID = uuid.NewV4().String()

	pd.VCM = vcm
	pd.colorPalette = palette
	pd.buffer = bytes.NewBuffer([]byte{})
	pd.img = image.NewRGBA(image.Rect(0, 0, int(dims[0]), int(dims[1])))

	pd.VCM.AddRenderer(pd.renderImage)

	pd.VCM.Model.SetTPS(tps)
	pd.VCM.SetRPS(tps)

	return pd
}
