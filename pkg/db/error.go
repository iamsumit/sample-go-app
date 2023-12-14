package db

import (
	"net/http"
	"strings"

	errpkg "github.com/iamsumit/sample-go-app/pkg/error"
)

// ErrorType is the error type for the errors returned by this package.
var ErrorType = "database"

var (
	ErrInternal = func(err error) *errpkg.Error {
		return NewError(
			err,
			http.StatusInternalServerError,
			errpkg.WithMessage("something went wrong internally"),
		)
	}
)

// NewError returns a new error instance with the given error message and
// status code.
//
// The attr values will be returned in the API response to the client.
func NewError(err error, status int, opts ...errpkg.Option) *errpkg.Error {
	// Set the status code and attribtues.
	opts = append(
		opts,
		errpkg.WithStatus(status),
	)

	// For any internal prefixed error, use a generic message.
	if strings.HasPrefix(err.Error(), "internal:") {
		opts = append(
			opts,
			errpkg.WithMessage("something went wrong internally"),
		)
	}

	e := errpkg.New(
		err,
		ErrorType,
		opts...,
	)

	return e
}
