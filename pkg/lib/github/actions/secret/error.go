package secret

import "errors"

// ErrMissingRepository is returned when no repository is provided.
var ErrMissingRepository = errors.New("the repository is required")
