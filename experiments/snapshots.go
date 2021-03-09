package experiments

import "image"

// TakeSnapshot take a snapshot.
func (pd *PetriDish) TakeSnapshot() image.Image {
	return pd.img
}

// SaveSnapshot save your snapshot in a file.
func (pd *PetriDish) SaveSnapshot(name string, format SnapshotImageFormat) error {
	return pd.storage.SaveSnapshot(name, pd.img, format)
}
