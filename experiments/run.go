package experiments

import "time"

// Ticks returns the current ticks in the model of your petri dish.
func (pd *PetriDish) Ticks() uint64 {
	return pd.ticks
}

// RunSync runs your petri dish in sync manner (await to finish).
func (pd *PetriDish) RunSync(duration time.Duration) {
	pd.VCM.Run(duration)
}

// Run ...
func (pd *PetriDish) Run(duration time.Duration) chan struct{} {
	done := make(chan struct{})

	go func() {
		pd.VCM.Run(duration)
		done <- struct{}{}
	}()

	return done
}

// RunSyncTicks runs your petri dish by n ticks.
func (pd *PetriDish) RunSyncTicks(ticks uint64) {
	pd.VCM.RunTicks(ticks)
}

// RunTicks runs your petri dish by ticks in async manner.
func (pd *PetriDish) RunTicks(ticks uint64) chan struct{} {
	done := make(chan struct{})

	go func() {
		pd.VCM.RunTicks(ticks)
		done <- struct{}{}
	}()

	return done
}
