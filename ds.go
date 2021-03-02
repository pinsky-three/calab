package calab

import uuid "github.com/satori/go.uuid"

// Space have the task to describe the lattice space of a any dynamical system.
type Space interface {
	Dims() []uint64
	State(i ...int64) uint64
	Neighbors(i ...int64) []uint64
}

// Evolvable saves how the space evolve over time.
type Evolvable interface {
	Evolve(space Space)
}

// DynamicalSystem represents a generalized dynamical system.
type DynamicalSystem struct {
	ID    string
	state Space
	rule  Evolvable

	ticks   uint64
	running bool
}

// BulkDynamicalSystem bulks a new DS.
func BulkDynamicalSystem(s Space, r Evolvable) *DynamicalSystem {
	return &DynamicalSystem{
		ID:      uuid.NewV4().String(),
		state:   s,
		rule:    r,
		ticks:   0,
		running: false,
	}
}

// State returns the state property of the DS structure.
func (ds *DynamicalSystem) State() Space {
	return ds.state
}

// type ComplexSystem interface {
// 	Run(*DynamicalSystem)
// }
