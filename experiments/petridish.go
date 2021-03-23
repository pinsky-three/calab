package experiments

import (
	"bytes"
	"image"

	"github.com/minskylab/calab"
)

type runMode string

const ticksMode runMode = "ticks"
const timeMode runMode = "time"

// PetriDish represents a fully instrumented system.
type PetriDish struct {
	ID             string
	VCM            *calab.VirtualComputationalModel
	buffer         *bytes.Buffer
	img            *image.RGBA
	colorPalette   calab.Palette
	dataHole       chan []byte
	ticks          uint64
	storage        Storage
	currentRunMode runMode
	Headless       bool
	timelapse      PetriDishTimelapse
}
