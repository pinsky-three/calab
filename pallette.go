package calab

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// Palette is a state to color mapping, also know as pallette.
type Palette map[uint64]color.Color

// NewPalette creates a new pallette
func NewPalette(colorStart colorful.Color, colorEnd colorful.Color, intervals int) Palette {
	palette := Palette{}

	for i := 0; i < intervals; i++ {
		palette[uint64(i)] = colorStart.BlendHsv(colorEnd, float64(i)/float64(intervals-1))
	}

	return palette
}

// NewCyclicPalette returns a cyclic palette.
func NewCyclicPalette(c1 colorful.Color, c2 colorful.Color, intervals int) Palette {
	palette := Palette{}

	for i := 0; i < intervals/2; i++ {
		palette[uint64(i)] = c1
		if intervals%2 != 0 {
			palette[uint64(i+1)] = c2
		}
	}

	return palette
}
