package lifelike

// New creates a new lifelike rule evolvable.
func New(r *Rule, bounder Bounder2D, neighborhood Neighborhood) (*LifeLike, error) {
	ll := &LifeLike{
		rule:         r,
		bounder:      bounder,
		neighborhood: neighborhood,
	}

	return ll, nil
}

// MustNew returns a new LifeLike rule, but if occurs some error it will be panic.
func MustNew(r *Rule, bounder Bounder2D, neighborhood Neighborhood) *LifeLike {
	ll, err := New(r, bounder, neighborhood)
	if err != nil {
		panic(err)
	}

	return ll
}
