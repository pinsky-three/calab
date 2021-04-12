package experiments

import (
	"sync"
	"time"
)

type ExperimentInterface struct {
	host string
	port int64
}

type Experiment struct {
	mu          sync.Locker
	dishes      map[string]*PetriDish
	dishesDones map[string]chan struct{}
	control     ExperimentInterface
}

func (exp *Experiment) syncDoneDish(dishID string, done chan struct{}) {
	exp.mu.Lock()
	exp.dishesDones[dishID] = done
	exp.mu.Unlock()

	go func(d chan struct{}, dishID string, mu sync.Locker) {
		<-d
		mu.Lock()
		delete(exp.dishesDones, dishID)
		mu.Unlock()
	}(done, dishID, exp.mu)

}

func (exp *Experiment) AddPetriDish(pd *PetriDish) {
	exp.dishes[pd.ID] = pd
}

func (exp *Experiment) DeletePetriDish(id string) {
	exp.mu.Lock()
	delete(exp.dishes, id)
	exp.mu.Unlock()
}

func (exp *Experiment) Run(dishID string, opts *Options) {
	var done chan struct{}

	if (opts.ticks != nil) == (opts.time != nil) {
		return
	}

	if opts.ticks != nil {
		done = exp.dishes[dishID].RunTicks(*opts.ticks)
	}

	if opts.time != nil {
		done = exp.dishes[dishID].Run(*opts.time)
	}

	exp.syncDoneDish(dishID, done)
}

func (exp *Experiment) Play(dishID string) {
	exp.dishes[dishID].Run(24 * time.Hour)
}

func (exp *Experiment) Pause(dishID string) {

}

func (exp *Experiment) Snapshot(dishID string) {

}

func (exp *Experiment) Timelapse(dishID string) {

}

/*

GET  https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}
POST https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/run?[time=[0-9]+[s|m|h]]&[ticks=[0-9]+]

POST https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/play
POST https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/pause
POST https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/snapshot/{snap-id}
GET  https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/timelapse
GET  https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/snapshots

GET https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/current[.png|.jpeg]
GET https://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/timelapse/{seq-id}[.png|.jpeg]

GET ws://ca.minsky.cc/experiments/{exp-id}/dish/{dish-id}/socket

*/
