package metrics

import "errors"

var (
	ErrMetersNotInitialized = errors.New("metrics not initialized")
	ErrInvalidType          = errors.New("invalid metrics type")
)
