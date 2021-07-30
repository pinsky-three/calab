package petridish

import "time"

// PetriDishTimelapse wraps intervals for create a timelapse
type PetriDishTimelapse struct {
	intervalTime  time.Duration
	intervalTicks uint64
}

// ProgramTimelapse program a time interval.
func (pd *PetriDish) ProgramTimelapse(interval time.Duration) {
	pd.timelapse.intervalTime = interval
	pd.currentRunMode = timeMode
}

// ProgramTickTimelapse programs a tick timelapse.
func (pd *PetriDish) ProgramTickTimelapse(interval uint64) {
	pd.timelapse.intervalTicks = interval
	pd.currentRunMode = ticksMode
}
