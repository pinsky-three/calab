package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/experiments/petridish"
	"github.com/minskylab/calab/experiments/utils"
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
	w, h := 1024, 1024

	// dynamic := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, lifelike.MooreNeighborhood(1, false))
	points := 4

	dynamic := voronoi.MustNew(points)

	initialState := map[uint64][][]int{}

	for i := 1; i < points+1; i++ {
		initialState[uint64(i)] = interpolate(
			horizontalLine(rand.Intn(w-3)+3, rand.Intn(h-3)+3),
		)
	}

	space := board.MustNew(w, h)
	space = space.Fill(board.FullState(0), dynamic)
	space = space.Fill(board.SpecificPositions(initialState), dynamic)

	ca := petridish.NewFromSystem(calab.BulkDynamicalSystem(space, dynamic))

	experiment := experiments.New()

	experiment.AddPetriDish(ca)

	experiment.Run(ca.ID, experiments.WithTicks(5000))

	time.Sleep(20 * time.Second)

	fmt.Printf("%s ends with %d ticks\n", ca.ID, ca.Ticks())

	utils.SaveSnapshotAsPNG(ca, fmt.Sprintf("%s.png", ca.ID))
}
