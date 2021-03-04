package renderers

import (
	"image"

	"github.com/minskylab/calab"
)

// ImageRenderer ...
type ImageRenderer interface {
	Render(n uint64, s calab.Space) image.Image
}
