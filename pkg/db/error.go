package db

import (
	"errors"
	"strings"
)

var (
	// ErrInternal is the default error message for any database related errors.
	//
	// Actual error will be logged.
	ErrInternal = errors.New("something went wrong internally")
)

// Error is a custom error type that holds the error message and status code.
//
// It can be validated by using errors.As() or err.(type) function.
type Error struct {
	// Err is the error message.
	Err error

	// Status is the HTTP status code.
	Status int

	// Attributes is a map of field name and error.
	Attributes map[string]interface{}
}

// NewError returns a new error instance with the given error message and
// status code.
//
// The attr values will be returned in the API response to the client.
func NewError(err error, status int, attr map[string]interface{}) *Error {
	return &Error{
		Err:        err,
		Status:     status,
		Attributes: attr,
	}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (e *Error) Error() string {
	msg := e.Err.Error()

	if strings.HasPrefix(msg, "internal:") {
		return ErrInternal.Error()
	}

	return msg
}
