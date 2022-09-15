package petridish

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
)

func (pd *PetriDish) SetPalette(newPalette calab.Palette) {
	pd.colorPalette = newPalette
}

func (pd *PetriDish) SetGradientPalette(colorStart colorful.Color, colorEnd colorful.Color) {
	states := int(pd.Model.System.Dynamic.Symbols())
	pd.SetPalette(calab.GradientPalette(colorStart, colorEnd, states))
}

func (pd *PetriDish) SetGradientWithVoidPalette(colorStart colorful.Color, colorEnd colorful.Color) {
	states := int(pd.Model.System.Dynamic.Symbols() - 1)

	finalPalette := map[uint64]color.Color{}
	finalPalette[0] = color.Black

	for i, p := range calab.GradientPalette(colorStart, colorEnd, states) {
		finalPalette[i+1] = p
	}

	pd.SetPalette(finalPalette)
}
