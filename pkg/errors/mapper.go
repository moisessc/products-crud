package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"products-crud/pkg/validator"
)

const (
	// UnmarshallErr identify an unmarshall error type
	UnmarshallErr ErrorType = "UNMARSHALL_ERROR"
	// ValidationErr identify a validation error type
	ValidationErr ErrorType = "VALIDATION_ERROR"
	// DomainErr identify a domain error type
	DomainErr ErrorType = "DOMAIN_ERROR"

	// internalServerErrorCode code to represent an internal server error
	internalServerErrorCode = "INTERNAL_SERVER_ERROR"
	// invalidRequestCode code to represent an invalid request
	invalidRequestCode = "INVALID_REQUEST"
)

// ErrorType type to specify an error type
type ErrorType string

// ApiResponse struct for error responses in the API
type ApiResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// MapError transform an error into custom error response
func MapError(err error, errType ErrorType) *ApiResponse {
	var msg, code string
	switch errType {
	case UnmarshallErr:
		msg = retrieveUnmarshalErrorMessage(err)
		code = invalidRequestCode
	case ValidationErr:
		msg = validator.RetrieveValidationErrorMessage(err)
		code = invalidRequestCode
	case DomainErr:
		msg = err.Error()
		code = retrieveDomainErrorCode(err)
	}

	return &ApiResponse{Message: msg, Code: code}
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
func retrieveDomainErrorCode(err error) string {
	switch {
	default:
		return internalServerErrorCode
	}
}
