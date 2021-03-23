package experiments

import "time"

type Options struct {
	time  *time.Duration
	ticks *uint64
}

func WithTime(duration string) *Options {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return &Options{}
	}

	return &Options{
		time:  &dur,
		ticks: nil,
	}
}

func WithTicks(ticks uint64) *Options {
	return &Options{
		ticks: &ticks,
		time:  nil,
	}
}
