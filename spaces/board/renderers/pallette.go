package renderers

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// Palette is a state to color mapping, also know as pallette.
type Palette map[uint64]color.Color

// NewPallette creates a new pallette
func NewPallette(colorStart colorful.Color, colorEnd colorful.Color, states int) Palette {
	pallette := Palette{}

	for i := 0; i < states; i++ {
		pallette[uint64(i)] = colorStart.BlendHcl(colorEnd, float64(i)/float64(states-1))
	}

	return pallette
}
