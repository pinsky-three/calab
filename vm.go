package calab

import (
	"time"
)

// VirtualMachine ...
type VirtualMachine struct {
	simulationDuration time.Duration
	model              *DynamicalSystem
	renderers          []Renderer
	RendersPerSecond   int
}

// NewVM ...
func NewVM(model *DynamicalSystem, rps int, renderers ...Renderer) *VirtualMachine {
	return &VirtualMachine{
		simulationDuration: time.Duration(0),
		RendersPerSecond:   rps,
		model:              model,
		renderers:          renderers,
	}
}

// Run ...
func (vm *VirtualMachine) Run(dt time.Duration) {
	ticks := make(chan uint64)
	done := make(chan struct{})
	lastTime := time.Now()

	vm.model.RunInfiniteSimulation(ticks, done)

	go func(done chan struct{}) {
		time.Sleep(dt)
		done <- struct{}{}
	}(done)

	go vm.model.Observe(ticks, func(n uint64, s Space) {
		// Limiting the renders per second.
		if time.Since(lastTime) < 1000/time.Duration(vm.RendersPerSecond)*time.Millisecond {
			return
		}

		// TODO: Update this rps limiter with an array of its, that's necessary for many renderers.
		for _, renderer := range vm.renderers {
			renderer(n, s)
		}

		lastTime = time.Now()
	})

}
