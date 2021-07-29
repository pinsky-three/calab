package experiments

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// SnapshotImageFormat ...
type SnapshotImageFormat string

// PNGFormat ...
const PNGFormat SnapshotImageFormat = "png"

// JPEGFormat ...
const JPEGFormat SnapshotImageFormat = "jpeg"

// Storage ...
type Storage interface {
	SavePetriDish(pd *PetriDish, withVM bool) error
	SaveSnapshot(filename string, img image.Image, format SnapshotImageFormat) error
}

// FileStorage implements a Experiment Storage.
type FileStorage struct {
	root string
	buff *bytes.Buffer
}

// SavePetriDish save petridish.
func (fs *FileStorage) SavePetriDish(pd *PetriDish, withVM bool) error {
	name := fmt.Sprintf("petri-%s.yaml", pd.Model.Model.ID[:6])

	filename := path.Join(fs.root, name)

	data, err := yaml.Marshal(pd)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = os.WriteFile(filename, data, 0644); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// SaveSnapshot saves a snapshot.
func (fs *FileStorage) SaveSnapshot(name string, img image.Image, format SnapshotImageFormat) error {
	buff := bytes.NewBuffer([]byte{})

	switch format {
	case PNGFormat:
		if err := png.Encode(buff, img); err != nil {
			return errors.WithStack(err)
		}
	case JPEGFormat:
		if err := jpeg.Encode(buff, img, &jpeg.Options{Quality: 100}); err != nil {
			return errors.WithStack(err)
		}
	default:
		return errors.New("invalid snapshot image format")
	}

	if err := os.WriteFile(name, buff.Bytes(), 0644); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// New returns a new file storage.
func New(root string) *FileStorage {
	return &FileStorage{
		root: root,
	}
}
