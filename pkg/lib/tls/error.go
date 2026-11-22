package tls

import "errors"

// ErrUnsupportedAlgorithm is returned when the key algorithm is not supported.
var ErrUnsupportedAlgorithm = errors.New("unsupported key algorithm")
