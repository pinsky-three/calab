package experiments

import (
	"sync"

	"github.com/minskylab/calab/experiments/petridish"
)

func New() *Experiment {
	exp := Experiment{}

	exp.mu = &sync.Mutex{}
	exp.dishes = map[string]*petridish.PetriDish{}
	exp.dishesDones = map[string]chan struct{}{}

	return &exp
}
