package renderers

import (
	"image"

	"github.com/minskylab/calab"
)

// BoardImageRenderer returns a new board image renderer.
type BoardImageRenderer struct {
	// Background color.Color

	image        *image.RGBA
	colorPalette Palette
}

// Render ...
func (bir *BoardImageRenderer) Render(n uint64, s calab.Space) image.Image {
	// board := s.(*board.Board2D)
	// w, h := len(board.Board), len(board.Board[0])
	dims := s.Dims()
	w, h := dims[0], dims[1]

	for i := int64(0); i < int64(w); i++ {
		for j := int64(0); j < int64(h); j++ {
			// bir.image.Set(i, j, bir.colorPalette[board[i][j]])
			bir.image.Set(int(i), int(j), bir.colorPalette[s.State(i, j)])
		}
	}

	return bir.image
}

// NewBoardImage ...
func NewBoardImage(w, h int, palette Palette) (*BoardImageRenderer, error) {
	return &BoardImageRenderer{
		image:        image.NewRGBA(image.Rect(0, 0, w, h)),
		colorPalette: palette,
	}, nil
}
