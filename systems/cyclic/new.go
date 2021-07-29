package cyclic

// NewRockPaperScissor returns a new Rock Paper Scissor.
func NewRockPaperScissor(bounder Bounder2D, neighborhood Neighborhood, totalStates, threshold int, stochastic int) (*RockPaperScissor, error) {
	return &RockPaperScissor{
		totalStates:  uint64(totalStates),
		threshold:    threshold,
		stochastic:   stochastic,
		bounder:      bounder,
		neighborhood: neighborhood,
	}, nil
}

// MustNewRockPaperScissor fails with panic occurs.
func MustNewRockPaperScissor(bounder Bounder2D, neighborhood Neighborhood, totalStates, threshold, stochastic int) *RockPaperScissor {
	ca, err := NewRockPaperScissor(bounder, neighborhood, threshold, totalStates, stochastic)
	if err != nil {
		panic(err)
	}

	return ca
}
