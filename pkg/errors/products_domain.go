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
	// ErrFailedToUpdateProduct is returned when some product could not be updated in the persistence layer
	ErrFailedToUpdateProduct = errors.New("product could not be updated")
	// ErrNothingToUpdate is returned when there is nothing to update in the persistence layer
	ErrNothingToUpdate = errors.New("the request do not have changes")
	// ErrFailedToDelete is returned when some product could not be deleted in the persistence layer
	ErrFailedToDeleteProduct = errors.New("product could not be deleted")
)
