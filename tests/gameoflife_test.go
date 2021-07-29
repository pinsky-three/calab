package tests

import (
	"testing"

	"github.com/minskylab/calab"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/lifelike"
)

var system128, system256, system512, system1024 *calab.DynamicalSystem

// var ticks chan uint64

func init() {
	classicMoore := lifelike.MooreNeighborhood(1, false)

	space128 := board.MustNew(128, 128)
	rule128 := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, classicMoore)
	system128 = calab.BulkDynamicalSystem(space128, rule128)

	space256 := board.MustNew(256, 256)
	rule256 := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, classicMoore)
	system256 = calab.BulkDynamicalSystem(space256, rule256)

	space512 := board.MustNew(512, 512)
	rule512 := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, classicMoore)
	system512 = calab.BulkDynamicalSystem(space512, rule512)

	space1024 := board.MustNew(1024, 1024)
	rule1024 := lifelike.MustNew(lifelike.GameOfLifeRule, lifelike.ToroidBounded, classicMoore)
	system1024 = calab.BulkDynamicalSystem(space1024, rule1024)
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

func benchmarkGoL1024xITicks(ticks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		system1024.RunSyncSimulation(uint64(ticks))
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

func BenchmarkGoL1024x1Ticks(b *testing.B)   { benchmarkGoL1024xITicks(1, b) }
func BenchmarkGoL1024x10Ticks(b *testing.B)  { benchmarkGoL1024xITicks(10, b) }
func BenchmarkGoL1024x20Ticks(b *testing.B)  { benchmarkGoL1024xITicks(20, b) }
func BenchmarkGoL1024x100Ticks(b *testing.B) { benchmarkGoL1024xITicks(100, b) }
func BenchmarkGoL1024x200Ticks(b *testing.B) { benchmarkGoL1024xITicks(200, b) }
func BenchmarkGoL1024x500Ticks(b *testing.B) { benchmarkGoL1024xITicks(500, b) }
