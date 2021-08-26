package petridish

import (
	"time"

	"github.com/minskylab/calab"
)

func NewTPSMonitor(pd *PetriDish) calab.Renderer {
	// TODO: Probably I'll need to add a mutex strategy here

	lastTime := time.Now()

	return func(n uint64, s calab.Space) {
		dur := time.Since(lastTime)
		pd.meanTPS = 1 / dur.Seconds()
		lastTime = time.Now()
	}
}

func WithTPSMonitor(pd *PetriDish) *PetriDish {
	pd.Model.AddRenderer(NewTPSMonitor(pd))

	return pd
}
