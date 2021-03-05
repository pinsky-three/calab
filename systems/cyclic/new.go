package cyclic

// NewRockPaperScissor returns a new Rock Paper Scissor.
func NewRockPaperScissor(states, threshold int, stochastic int) (*RockPaperScissor, error) {
	countsMap := map[uint64]int{}

	for i := 0; i < states; i++ {
		countsMap[uint64(i)] = 0
	}

	return &RockPaperScissor{
		countsMap:  countsMap,
		threshold:  threshold,
		stochastic: stochastic,
	}, nil
}

// MustNewRockPaperScissor fails is panic occurs.
func MustNewRockPaperScissor(states, threshold int, stochastic int) *RockPaperScissor {
	ca, err := NewRockPaperScissor(states, threshold, stochastic)
	if err != nil {
		panic(err)
	}

	return ca
}
