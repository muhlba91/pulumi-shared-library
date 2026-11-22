package file

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// Hash computes the hex-encoded SHA-512 hash of the file at the given path.
// path: The path to the file to hash.
func Hash(path string) (*string, error) {
	data, err := ReadContents(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %q for hashing: %w", path, err)
	}

	sum := sha512.Sum512([]byte(data))
	hash := hex.EncodeToString(sum[:])
	return &hash, nil
}
