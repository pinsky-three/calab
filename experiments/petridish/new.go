package petridish

import (
	"image"

	"github.com/minskylab/calab"
	uuid "github.com/satori/go.uuid"
)

type PetriDishOption func(*PetriDish) *PetriDish

func NewFromSpaceAndDynamic(space calab.Space, dynamic calab.Dynamic, opts ...PetriDishOption) *PetriDish {
	return NewFromSystem(calab.BulkDynamicalSystem(space, dynamic), opts...)
}

func NewFromSystem(system *calab.DynamicalSystem, opts ...PetriDishOption) *PetriDish {
	vm := calab.NewVM(system)
	return NewFromVCM(vm, opts...)
}

func NewDefault(system *calab.DynamicalSystem) *PetriDish {
	return NewFromSystem(system, WithTPSMonitor)
}

// NewFromVCM returns a new petridish.
func NewFromVCM(vcm *calab.VirtualComputationalModel, opts ...PetriDishOption) *PetriDish {
	dims := vcm.System.Space.Dims()

	pd := &PetriDish{}
	pd.ID = uuid.NewV4().String()
	pd.observers = []chan image.Image{}
	pd.cache = map[uint64]image.Image{}
	pd.Model = vcm
	pd.colorPalette = calab.MonochromePalette(vcm.System.Dynamic.Symbols())
	pd.img = image.NewRGBA(image.Rect(0, 0, int(dims[0]), int(dims[1])))

	pd.Model.AddRenderer(pd.observation)

	for _, opt := range opts {
		pd = opt(pd)
	}

	// pd.Model.System.SetTPS(DefaultTPS)
	// pd.Model.SetRPS(DefaultTPS)

	return pd
}
