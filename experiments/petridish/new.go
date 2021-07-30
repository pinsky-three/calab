package petridish

import (
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
	dims := vcm.System.Space.Dims()

	pd := &PetriDish{}
	pd.ID = uuid.NewV4().String()
	pd.observers = []chan image.Image{}
	pd.cache = map[uint64]image.Image{}
	pd.Model = vcm
	pd.colorPalette = calab.MonochromePalette(vcm.System.Dynamic.Symbols())
	pd.img = image.NewRGBA(image.Rect(0, 0, int(dims[0]), int(dims[1])))

	pd.Model.AddRenderer(pd.observation)

	pd.Model.System.SetTPS(DefaultTPS)
	pd.Model.SetRPS(DefaultTPS)

	return pd
}
