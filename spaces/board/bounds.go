package board

// Bounder represents a new bounder mapper
type Bounder func(w, h, x, y int64) (int64, int64)

// ToroidBounded is a toroind kind bounder
func ToroidBounded(w, h, x, y int64) (int64, int64) {
	if x > int64(w-1) {
		x = 0
	}

	if x < 0 {
		x = int64(w - 1)
	}

	if y > int64(h-1) {
		y = 0
	}

	if y < 0 {
		y = int64(h - 1)
	}

	return x, y
}
