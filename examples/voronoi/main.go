package main

import (
	"math/rand"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/voronoi"
)

func horizontalLine(w, h int) [][]int {
	return [][]int{
		{w - 3, h},
		{w - 2, h},
		{w - 1, h},
		{w, h},
		{w + 1, h},
		{w + 2, h},
		{w + 3, h},
	}
}

func interpolate(cloud ...[][]int) [][]int {
	result := [][]int{}

	for _, cl := range cloud {
		result = append(result, cl...)
	}

	return result
}

func main() {
	w, h := 2048, 2048

	points := 32

	dynamic := voronoi.MustNew(points)

	initialState := map[uint64][][]int{}

	for i := 1; i < points+1; i++ {
		initialState[uint64(i)] = horizontalLine(rand.Intn(w-3)+3, rand.Intn(h-3)+3)
	}

	space := board.MustNew(w, h)
	space = space.Fill(board.FullState(0), dynamic)
	space = space.Fill(board.SpecificPositions(initialState), dynamic)

	ca := petridish.NewFromSystem(calab.BulkDynamicalSystem(space, dynamic), petridish.WithTPSMonitor)

	experiment := experiments.New()
	experiment.AddPetriDish(ca)

	timelapseOptions := &experiments.TimeLapseOptions{
		Debug:       true,
		DeleteAfter: false,
	}

	done := ca.RunTicks(1000)
	if err := experiment.Timelapse(ca.ID, done, timelapseOptions); err != nil {
		panic(err)
	}
}
