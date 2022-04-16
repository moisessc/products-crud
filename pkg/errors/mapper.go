package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"products-crud/pkg/validator"
)

const (
	// UnmarshallErr identify an unmarshall error type
	UnmarshallErr ErrorType = "UNMARSHALL_ERROR"
	// InvalidPathParam identify an path param error type
	InvalidPathParam ErrorType = "PATH_PARAM_PARSE_ERROR"
	// ValidationErr identify a validation error type
	ValidationErr ErrorType = "VALIDATION_ERROR"
	// DomainErr identify a domain error type
	DomainErr ErrorType = "DOMAIN_ERROR"

	// internalServerErrorCode code to represent an internal server error
	internalServerErrorCode = "INTERNAL_SERVER_ERROR"
	// invalidRequestCode code to represent an invalid request
	invalidRequestCode = "INVALID_REQUEST"
	// notFoundCode code to represent that the resource could not be found
	notFoundCode = "NOT_FOUND"
	// nothingToUpdate code to represent that there is nothing to update
	nothingToUpdate = "NOTHING_TO_UPDATE"
)

// ErrorType type to specify an error type
type ErrorType string

// ApiResponse struct for error responses in the API
type ApiResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// MapError transform an error into custom error response
func MapError(err error, errType ErrorType) (*ApiResponse, int) {
	var msg, code string
	var statusCode int
	switch errType {
	case UnmarshallErr:
		msg = retrieveUnmarshalErrorMessage(err)
		code = invalidRequestCode
		statusCode = http.StatusBadRequest
	case InvalidPathParam:
		msg = "invalid id"
		code = invalidRequestCode
		statusCode = http.StatusBadRequest
	case ValidationErr:
		msg = validator.RetrieveValidationErrorMessage(err)
		code = invalidRequestCode
		statusCode = http.StatusBadRequest
	case DomainErr:
		msg = err.Error()
		code, statusCode = retrieveDomainErrorCode(err)
	}

	return &ApiResponse{Message: msg, Code: code}, statusCode
}

// retrieveUnmarshalErrorInformation retrieves the information when the bind method fails
func retrieveUnmarshalErrorMessage(err error) string {
	var field, expected, got string
	var jsErr *json.UnmarshalTypeError
	if errors.As(err, &jsErr) {
		field = jsErr.Field
		expected = jsErr.Type.Name()
		got = jsErr.Value
	}

	if strings.Contains(expected, "int") || strings.Contains(expected, "float") {
		expected = "number"
	}
	return fmt.Sprint("unmarshal error data type, got: ", got, ", expected: ", expected, " in ", field, " param")
}

// retrieveDomainErrorCode retrieves the error code of one domain error
func retrieveDomainErrorCode(err error) (string, int) {
	switch {
	case errors.Is(err, ErrProductNotFound):
		return notFoundCode, http.StatusNotFound
	case errors.Is(err, ErrNothingToUpdate):
		return nothingToUpdate, http.StatusBadRequest
	default:
		return internalServerErrorCode, http.StatusInternalServerError
	}
}
