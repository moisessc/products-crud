package errors

import "errors"

var (
	// ErrFailedToSaveProduct is returned when saving a product fails in the persistence layer
	ErrFailedToSaveProduct = errors.New("product could not be saved")
)
