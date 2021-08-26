package calab

import (
	"time"
)

// Tick execute one dynamical tick system evolution.
func (ds *DynamicalSystem) Tick() {
	elapsedTime := time.Since(ds.lastTime)
	expectedDuration := 1000 / time.Duration(ds.ticksPerSecond) * time.Millisecond

	if elapsedTime < expectedDuration {
		// fmt.Println("waiting", expectedDuration-elapsedTime)
		time.Sleep(expectedDuration - elapsedTime)
	}

	ds.Dynamic.Evolve(ds.Space)

	ds.lastTime = time.Now()
	ds.ticks++
}

// RunSimulation run simulation 'ticks' ticks.
func (ds *DynamicalSystem) RunSimulation(ticks uint64, cn chan uint64) {
	ds.running = true
	go func(cn chan uint64, running *bool) {
		for i := uint64(0); i < ticks; i++ {
			ds.Tick()
			cn <- ds.ticks
		}
		close(cn)
		*running = false
	}(cn, &ds.running)
}

// RunSyncSimulation executes a 'ticks' ticks synchornous simulation.
func (ds *DynamicalSystem) RunSyncSimulation(ticks uint64) {
	ds.running = true
	for i := uint64(0); i < ticks; i++ {
		ds.Tick()
	}
	ds.running = false
}

// RunInfiniteSimulation runs a infinite (but closable) simulation.
func (ds *DynamicalSystem) RunInfiniteSimulation(cn chan uint64, finish chan struct{}) {
	ds.running = true

	go func(f chan struct{}, running *bool) {
		<-finish
		*running = false
	}(finish, &ds.running)

	go func(cn chan uint64, running *bool) {
		cn <- ds.ticks
		for *running {
			ds.Tick()
			cn <- ds.ticks
		}
		close(cn)
		close(finish)
	}(cn, &ds.running)
}

func (ds *DynamicalSystem) Pause() {
	ds.running = false
}

// Observe execute a function on every tick from ticker channel.
func (ds *DynamicalSystem) Observe(cn chan uint64, observer ObserverFunction) {
	for n := range cn {
		observer(n, ds.Space)
	}
}
