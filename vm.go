package calab

import (
	"time"
)

// VirtualComputationalModel ...
type VirtualComputationalModel struct {
	Model            *DynamicalSystem
	rendersPerSecond int
	renderers        []Renderer
}

// NewVM ...
func NewVM(model *DynamicalSystem, renderers ...Renderer) *VirtualComputationalModel {
	return &VirtualComputationalModel{
		rendersPerSecond: 60,
		Model:            model,
		renderers:        renderers,
	}
}

// AddRenderer ...
func (vm *VirtualComputationalModel) AddRenderer(r Renderer) {
	vm.renderers = append(vm.renderers, r)
}

// SetRPS sets renders per second rate.
func (vm *VirtualComputationalModel) SetRPS(rendersPerSecond int) {
	vm.rendersPerSecond = rendersPerSecond
}

// Run ...
func (vm *VirtualComputationalModel) Run(dt time.Duration) {
	ticks := make(chan uint64)
	done := make(chan struct{})
	lastTime := time.Now()

	go func(done chan struct{}) {
		time.Sleep(dt)
		done <- struct{}{}
	}(done)

	vm.Model.RunInfiniteSimulation(ticks, done)

	vm.Model.Observe(ticks, func(n uint64, s Space) {
		// Limiting the renders per second.
		if time.Since(lastTime) < 1000/time.Duration(vm.rendersPerSecond)*time.Millisecond {
			return
		}

		// TODO: Update this rps limiter with an array of its, that's necessary for many renderers.
		for _, renderer := range vm.renderers {
			renderer(n, s)
		}

		lastTime = time.Now()
	})
}

// RunTicks runs your simulation for n ticks.
func (vm *VirtualComputationalModel) RunTicks(ticks uint64) {
	ticksChannel := make(chan uint64)

	vm.Model.RunSimulation(ticks, ticksChannel)

	vm.Model.Observe(ticksChannel, func(n uint64, s Space) {
		for _, renderer := range vm.renderers {
			renderer(n, s)
		}
	})
}
