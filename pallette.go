package calab

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// Palette is a state to color mapping, also know as pallette.
type Palette map[uint64]color.Color

// GradientPalette creates a new pallette
func GradientPalette(colorStart colorful.Color, colorEnd colorful.Color, intervals int) Palette {
	palette := Palette{}

	for i := 0; i < intervals; i++ {
		palette[uint64(i)] = colorStart.BlendHsv(colorEnd, float64(i)/float64(intervals-1))
	}

	return palette
}

// CyclicPalette returns a cyclic palette.
func CyclicPalette(c1 colorful.Color, c2 colorful.Color, intervals int) Palette {
	palette := Palette{}

	for i := 0; i < intervals/2; i++ {
		palette[uint64(i)] = c1
		palette[uint64(i+1)] = c2
	}

	if intervals%2 != 0 {
		palette[uint64(intervals/2+1)] = c1
	}

	return palette
}

func MonochromePalette(states uint64) Palette {
	palette := Palette{}

	for i := 0; i < int(states); i++ {
		grey := uint8(i * 255 / (int(states) - 1))
		palette[uint64(i)] = color.RGBA{grey, grey, grey, 255}
	}

	return palette
}
