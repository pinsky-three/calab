package petridish

import (
	"bytes"
	"image"

	"github.com/minskylab/calab"
)

type runMode string

const (
	ticksMode runMode = "ticks"
	timeMode  runMode = "time"
)

// PetriDish represents a fully instrumented system.
type PetriDish struct {
	ID             string
	Model          *calab.VirtualComputationalModel
	buffer         *bytes.Buffer
	img            *image.RGBA
	colorPalette   calab.Palette
	dataHole       chan []byte
	ticks          uint64
	storage        Storage
	currentRunMode runMode
	timelapse      PetriDishTimelapse
}
