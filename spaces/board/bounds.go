package board

// Bounder2D represents a new bounder mapper
type Bounder2D func(w, h, xi, yi int64) (int64, int64)

// ToroidBounded is a toroind kind bounder
func ToroidBounded(w, h, xi, yi int64) (int64, int64) {
	if xi > int64(w-1) {
		xi = 0
	}

	if xi < 0 {
		xi = int64(w - 1)
	}

	if yi > int64(h-1) {
		yi = 0
	}

	if yi < 0 {
		yi = int64(h - 1)
	}

	return xi, yi
}
