package validator

import (
	errpkg "github.com/iamsumit/sample-go-app/pkg/error"
)

// ErrorType is the error type for the errors returned by this package.
var ErrorType = "validator"

// NewError returns a new error instance with the given error message and
// status code.
//
// The attr values will be returned in the API response to the client.
func NewError(err error, status int, attr map[string]interface{}, opts ...errpkg.Option) *errpkg.Error {
	// Set the status code and attribtues.
	opts = append(
		opts,
		errpkg.WithStatus(status),
		errpkg.WithAttributes(attr),
	)

	e := errpkg.New(
		err,
		ErrorType,
		opts...,
	)

	return e
}
