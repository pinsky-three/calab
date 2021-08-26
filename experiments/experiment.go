package experiments

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path"
	"sync"
	"time"

	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/experiments/utils"
	"github.com/pkg/errors"
)

// type ExperimentInterface struct {
// 	host string
// 	port int64
// }

type Experiment struct {
	mu          sync.Locker
	dishes      map[string]*petridish.PetriDish
	dishesDones map[string]chan struct{}
	// control     ExperimentInterface
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

func (exp *Experiment) AddPetriDish(pd *petridish.PetriDish) {
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

	if _, exist := exp.dishes[dishID]; !exist {
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

// func (exp *Experiment) Play(dishID string) {
// 	exp.dishes[dishID].Run(24 * time.Hour)
// }

// func (exp *Experiment) Pause(dishID string) {
// }

func (exp *Experiment) Snapshot(dishID string) (string, uint64) {
	filename := "snapshot_" + dishID + "_" + time.Now().Format("2006_01_02_15_04_05") + ".png"
	if err := utils.SaveSnapshotAsPNG(exp.dishes[dishID], filename); err != nil {
		panic(err)
	}

	return filename, exp.dishes[dishID].Ticks()
}

func (exp *Experiment) Ticks(dishID string) uint64 {
	return exp.dishes[dishID].Ticks()
}

// type ObserveFrame struct {
// 	image image.Image
// 	tick  uint64
// 	tps   float64
// }

func (exp *Experiment) Observe(dishID string) (chan image.Image, error) {
	channel := make(chan image.Image, 1)
	exp.dishes[dishID].RegisterNewObserver(channel)

	return channel, nil
}

type TimeLapseOptions struct {
	OutputFilename string
	Debug          bool
	DeleteAfter    bool
}

func (exp *Experiment) Timelapse(dishID string, done chan struct{}, opts *TimeLapseOptions) error {
	ca := exp.dishes[dishID]

	tempFolder := path.Join(os.TempDir(), "calab")
	imagesPath := path.Join(tempFolder, ca.ID)

	frames, err := exp.Observe(ca.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := os.MkdirAll(imagesPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	go func() {
		for frame := range frames {
			// fmt.Printf("frame: %d\n", gameOfLife.Ticks())

			filename := fmt.Sprintf("frame-%d.png", ca.Ticks())
			filepath := path.Join(imagesPath, filename)

			go func(filepath string, image image.Image) {
				f, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					panic(err)
				}

				png.Encode(f, image)
				f.Close()
			}(filepath, frame)
		}
	}()

	<-done

	outVideo := fmt.Sprintf("%s.mp4", ca.ID)
	if opts != nil && opts.OutputFilename != "" {
		outVideo = opts.OutputFilename
	}

	// TODO: Add more function parameters to specify ffmpeg modifiers.

	if err := exec.Command("ffmpeg",
		"-framerate", "24",
		"-i", path.Join(imagesPath, "frame-%d.png"),
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		"-crf", "25",
		"-preset", "slow",
		"-tune", "animation",
		"-f", "mp4",
		"-y",
		outVideo).Run(); err != nil {
		return errors.WithStack(err)
	}

	if opts == nil || (opts != nil && opts.Debug) {
		fmt.Printf("id: %s\n", ca.ID)
		fmt.Printf("frames: %s\n", imagesPath)
		fmt.Printf("video: %s\n", outVideo)
		fmt.Printf("total ticks: %d\n", ca.Ticks())
		fmt.Printf("mean tps: %.3f\n", ca.GetMeanTPS())
	}

	if opts != nil && opts.DeleteAfter {
		if err := os.RemoveAll(imagesPath); err != nil {
			return errors.WithStack(err)
		}
	}

	close(frames)
	close(done)

	return nil
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
