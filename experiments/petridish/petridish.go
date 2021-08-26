package petridish

import (
	"image"

	"github.com/minskylab/calab"
)

type runMode string

const (
	ticksMode runMode = "ticks"
	timeMode  runMode = "time"
)

// PetriDish is a representation of a virtual computation model.
// It is a RGBA Image interpretation from an Space.
type PetriDish struct {
	ID    string
	Model *calab.VirtualComputationalModel

	img            *image.RGBA
	colorPalette   calab.Palette
	ticks          uint64
	currentRunMode runMode
	timelapse      PetriDishTimelapse
	observers      []chan image.Image
	cache          map[uint64]image.Image
	meanTPS        float64
}
