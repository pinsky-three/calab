package experiments

import "image"

// TakeSnapshot take a snapshot.
func (pd *PetriDish) TakeSnapshot() image.Image {
	return pd.img
}
