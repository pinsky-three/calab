package voronoi

type Voronoi struct {
	stateToIndex map[int]int
	indexToState map[int]int
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
	primes := sieveOfEratosthenes(10_000, totalPoints)

	stateToIndex := map[int]int{}
	indexToState := map[int]int{}

	stateToIndex[0] = 0
	indexToState[0] = 0

	for i, p := range primes {
		stateToIndex[p] = i + 1
		indexToState[i+1] = p
	}

	return &Voronoi{
		stateToIndex: stateToIndex,
		indexToState: indexToState,
	}, nil
}

func MustNew(totalStates int) *Voronoi {
	v, err := New(totalStates)
	if err != nil {
		panic(err)
	}

	return v
}
