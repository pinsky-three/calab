package calab

import "time"

// VirtualMachine ...
type VirtualMachine struct {
	model     *DynamicalSystem
	renderers []Renderer
}

// NewVM ...
func NewVM(model *DynamicalSystem, renderers ...Renderer) *VirtualMachine {
	return &VirtualMachine{
		model:     model,
		renderers: renderers,
	}
}

// Run ...
func (vm *VirtualMachine) Run(dt time.Duration) {
	ticks := make(chan uint64)
	done := make(chan struct{})

	vm.model.RunInfiniteSimulation(ticks, done)

	go func(done chan struct{}) {
		time.Sleep(dt)
		done <- struct{}{}
	}(done)

	vm.model.Observe(ticks, func(n uint64, s Space) {
		for _, renderer := range vm.renderers {
			renderer(n, s)
		}
	})

}
