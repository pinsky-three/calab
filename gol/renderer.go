package gol

import (
	"image"

	"github.com/minskylab/calab"
)

type ImageRenderer struct {
	reusableImage *image.Gray
	images        chan image.Image
}

func (ir *ImageRenderer) Render(n uint64, s calab.Space) {
	ca := s.(*CA)

	flat := []byte{}
	for _, r := range ca.board {
		flat = append(flat, r...)
	}

	// w := len(ca.board)
	// h := len(ca.board[0])

	for i := range flat {
		ir.reusableImage.Pix[i] = flat[i] * 255
	}

	ir.images <- ir.reusableImage
}

func NewImageRenderer(images chan image.Image, w, h int) *ImageRenderer {
	return &ImageRenderer{
		images:        images,
		reusableImage: image.NewGray(image.Rect(0, 0, w, h)),
	}
}
