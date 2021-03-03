package tests

import (
	"testing"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
)

var system128, system256, system512, system820 *calab.DynamicalSystem

// var ticks chan uint64

func init() {
	nh128 := board.MooreNeighborhood(1, false)
	bound128 := board.ToroidBounded()
	space128 := board.MustNew(128, 128, nh128, bound128, board.RandomInit, board.UniformNoise)
	rule128 := lifelike.MustNew(lifelike.GameOfLifeRule)
	system128 = calab.BulkDynamicalSystem(space128, rule128)

	nh256 := board.MooreNeighborhood(1, false)
	bound256 := board.ToroidBounded()
	space256 := board.MustNew(256, 256, nh256, bound256, board.RandomInit, board.UniformNoise)
	rule256 := lifelike.MustNew(lifelike.GameOfLifeRule)
	system256 = calab.BulkDynamicalSystem(space256, rule256)

	nh512 := board.MooreNeighborhood(1, false)
	bound512 := board.ToroidBounded()
	space512 := board.MustNew(512, 512, nh512, bound512, board.RandomInit, board.UniformNoise)
	rule512 := lifelike.MustNew(lifelike.GameOfLifeRule)
	system512 = calab.BulkDynamicalSystem(space512, rule512)

	nh820 := board.MooreNeighborhood(1, false)
	bound820 := board.ToroidBounded()
	space820 := board.MustNew(820, 820, nh820, bound820, board.RandomInit, board.UniformNoise)
	rule820 := lifelike.MustNew(lifelike.GameOfLifeRule)
	system820 = calab.BulkDynamicalSystem(space820, rule820)
}

func benchmarkGoL128xITicks(ticks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		system128.RunSyncSimulation(uint64(ticks))
	}
}

func benchmarkGoL256xITicks(ticks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		system256.RunSyncSimulation(uint64(ticks))
	}
}

func benchmarkGoL512xITicks(ticks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		system512.RunSyncSimulation(uint64(ticks))
	}
}

func benchmarkGoL820xITicks(ticks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		system512.RunSyncSimulation(uint64(ticks))
	}
}

func BenchmarkGoL128x1Ticks(b *testing.B)   { benchmarkGoL128xITicks(1, b) }
func BenchmarkGoL128x10Ticks(b *testing.B)  { benchmarkGoL128xITicks(10, b) }
func BenchmarkGoL128x20Ticks(b *testing.B)  { benchmarkGoL128xITicks(20, b) }
func BenchmarkGoL128x100Ticks(b *testing.B) { benchmarkGoL128xITicks(100, b) }
func BenchmarkGoL128x200Ticks(b *testing.B) { benchmarkGoL128xITicks(200, b) }
func BenchmarkGoL128x500Ticks(b *testing.B) { benchmarkGoL128xITicks(500, b) }

func BenchmarkGoL256x1Ticks(b *testing.B)   { benchmarkGoL256xITicks(1, b) }
func BenchmarkGoL256x10Ticks(b *testing.B)  { benchmarkGoL256xITicks(10, b) }
func BenchmarkGoL256x20Ticks(b *testing.B)  { benchmarkGoL256xITicks(20, b) }
func BenchmarkGoL256x100Ticks(b *testing.B) { benchmarkGoL256xITicks(100, b) }
func BenchmarkGoL256x200Ticks(b *testing.B) { benchmarkGoL256xITicks(200, b) }
func BenchmarkGoL256x500Ticks(b *testing.B) { benchmarkGoL256xITicks(500, b) }

func BenchmarkGoL512x1Ticks(b *testing.B)   { benchmarkGoL512xITicks(1, b) }
func BenchmarkGoL512x10Ticks(b *testing.B)  { benchmarkGoL512xITicks(10, b) }
func BenchmarkGoL512x20Ticks(b *testing.B)  { benchmarkGoL512xITicks(20, b) }
func BenchmarkGoL512x100Ticks(b *testing.B) { benchmarkGoL512xITicks(100, b) }
func BenchmarkGoL512x200Ticks(b *testing.B) { benchmarkGoL512xITicks(200, b) }
func BenchmarkGoL512x500Ticks(b *testing.B) { benchmarkGoL512xITicks(500, b) }

func BenchmarkGoL820x1Ticks(b *testing.B)   { benchmarkGoL820xITicks(1, b) }
func BenchmarkGoL820x10Ticks(b *testing.B)  { benchmarkGoL820xITicks(10, b) }
func BenchmarkGoL820x20Ticks(b *testing.B)  { benchmarkGoL820xITicks(20, b) }
func BenchmarkGoL820x100Ticks(b *testing.B) { benchmarkGoL820xITicks(100, b) }
func BenchmarkGoL820x200Ticks(b *testing.B) { benchmarkGoL820xITicks(200, b) }
func BenchmarkGoL820x500Ticks(b *testing.B) { benchmarkGoL820xITicks(500, b) }
