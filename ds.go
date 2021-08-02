package calab

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const DefaultTPS = 1 << 20

// Space have the task to describe the lattice space of a any dynamical system.
type Space interface {
	Dims() []uint64
	Space() []uint64
	State(i ...int64) uint64

	Branch(space []uint64) Space
}

// Evolvable saves how the space evolve over time.
type Evolvable interface {
	Evolve(space Space) Space
	Symbols() uint64
}

// DynamicalSystem represents a generalized dynamical system.
type DynamicalSystem struct {
	ID      string
	Space   Space
	Dynamic Evolvable

	ticksPerSecond int
	ticks          uint64

	running  bool
	lastTime time.Time
}

// BulkDynamicalSystem bulks a new DS.
func BulkDynamicalSystem(s Space, d Evolvable) *DynamicalSystem {
	return &DynamicalSystem{
		ID:             uuid.NewV4().String(),
		Space:          s,
		Dynamic:        d,
		ticks:          0,
		running:        false,
		ticksPerSecond: DefaultTPS,
	}
}

// SetTPS set ticks per second for your dynamical model.
func (ds *DynamicalSystem) SetTPS(ticksPerSecond int) {
	ds.ticksPerSecond = ticksPerSecond
}
