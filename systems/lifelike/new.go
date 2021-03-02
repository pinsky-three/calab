package lifelike

// New creates a new lifelike rule evolvable.
func New(r *Rule) (*LifeLike, error) {
	ll := &LifeLike{
		rule: r,
	}

	return ll, nil
}

// MustNew returns a new LifeLike rule, but if occurs some error it will be panic.
func MustNew(r *Rule) *LifeLike {
	ll, err := New(r)
	if err != nil {
		panic(err)
	}

	return ll
}
