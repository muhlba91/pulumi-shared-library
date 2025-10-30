package dir

import (
	"os"
)

// Create creates the directory at the given path.
// Attention: removes the entire directory beforehand!
// path: directory path to create.
func Create(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return os.MkdirAll(path, 0o750)
}
