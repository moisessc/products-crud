package errors

import "errors"

var (
	// ErrFailedToSaveProduct is returned when saving a product fails in the persistence layer
	ErrFailedToSaveProduct = errors.New("product could not be saved")
	// ErrFailedToRetrieveProducts is returned when the products could not be retrieved in the persistence layer
	ErrFailedToRetrieveProducts = errors.New("products could not be retrieved")
)
