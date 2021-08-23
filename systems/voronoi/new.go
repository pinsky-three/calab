package voronoi

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func New(totalStates int) (*Voronoi, error) {
	return &Voronoi{
		totalStates: totalStates,
		states:      sieveOfEratosthenes(totalStates),
	}, nil
}

func MustNew(totalStates int) *Voronoi {
	v, err := New(totalStates)
	if err != nil {
		panic(err)
	}

	return v
}
