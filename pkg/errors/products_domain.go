package errors

import "errors"

var (
	// ErrFailedToSaveProduct is returned when saving a product fails in the persistence layer
	ErrFailedToSaveProduct = errors.New("product could not be saved")
	// ErrFailedToRetrieveProducts is returned when the products could not be retrieved in the persistence layer
	ErrFailedToRetrieveProducts = errors.New("products could not be retrieved")
	// ErrProductNotFound is returned when some product is not found in the persistence layer
	ErrProductNotFound = errors.New("product could not be found")
	// ErrFailedToRetrieveProduct is returned when some product could not be retrieved in the persistence layer
	ErrFailedToRetrieveProduct = errors.New("product could not be retrieved")
)
