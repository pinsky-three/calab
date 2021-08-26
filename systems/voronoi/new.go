package voronoi

type Voronoi struct {
	// stateToIndex map[int]uint64
	// indexToState map[uint64]int
	totalPoints int
}

func sieveOfEratosthenes(N int, expectedPrimes int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] {
			continue
		}

		primes = append(primes, i)

		if len(primes) == expectedPrimes {
			return
		}

		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}

	return
}

func New(totalPoints int) (*Voronoi, error) {
	// primes := sieveOfEratosthenes(10_000, totalPoints)

	// stateToIndex := map[int]uint64{}
	// indexToState := map[uint64]int{}

	// stateToIndex[0] = 0
	// indexToState[0] = 0

	// for i, p := range primes {
	// 	ind := uint64(i + 1)

	// 	stateToIndex[p] = ind
	// 	indexToState[ind] = p
	// }

	return &Voronoi{
		// stateToIndex: stateToIndex,
		// indexToState: indexToState,
		totalPoints: totalPoints,
	}, nil
}

func MustNew(totalStates int) *Voronoi {
	v, err := New(totalStates)
	if err != nil {
		panic(err)
	}

	return v
}
