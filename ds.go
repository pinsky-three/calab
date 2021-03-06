package calab

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

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
	ID             string
	ticksPerSecond int

	Space Space
	rule  Evolvable

	ticks    uint64
	running  bool
	lastTime time.Time
}

// BulkDynamicalSystem bulks a new DS.
func BulkDynamicalSystem(s Space, r Evolvable, tps int) *DynamicalSystem {
	return &DynamicalSystem{
		ID:             uuid.NewV4().String(),
		Space:          s,
		rule:           r,
		ticks:          0,
		running:        false,
		ticksPerSecond: tps,
	}
}

// SetTPS set ticks per second for your dynamical model.
func (ds *DynamicalSystem) SetTPS(ticksPerSecond int) {
	ds.ticksPerSecond = ticksPerSecond
}
