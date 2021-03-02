package cyclic

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

func mustHexColor(val string) colorful.Color {
	s, err := colorful.Hex(val)
	if err != nil {
		panic(err)
	}

	return s
}

// NewRockPaperScissor ...
func NewRockPaperScissor(w, h int, states int, threshold int, stocasticity int, randomSeed int64, images chan image.Image) (*PaperRockSissor, error) {
	rand.Seed(randomSeed)

	c1 := mustHexColor("#fcd8ba")
	c2 := mustHexColor("#faad65")

	colorMap := map[byte]color.Color{}

	colorMap[0] = c1

	for i := 1; i < states; i++ {
		colorMap[byte(i)] = c1.BlendHcl(c2, float64(i+1)/float64(states))
	}

	countsMap := map[byte]int{}

	for c := range colorMap {
		countsMap[c] = 0
	}

	board := make([][]byte, h)
	for i := range board {
		board[i] = make([]byte, w)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = byte(rand.Intn(len(colorMap)))
		}
	}

	return &PaperRockSissor{
		board:        board,
		countsMap:    countsMap,
		threshold:    threshold,
		frame:        image.NewRGBA(image.Rect(0, 0, w, h)),
		images:       images,
		colormap:     colorMap,
		stocasticity: stocasticity,
	}, nil
}
