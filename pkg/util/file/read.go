package file

import (
	"os"
)

// ReadContents reads and returns the contents of the file at path.
// path: file path to read from.
func ReadContents(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
