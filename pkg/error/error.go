package error

import "net/http"

// Option type to use the options pattern to set the configuration.
type Option func(*Error)

// Error type to use the error pattern to return the error.
type Error struct {
	// Type of the error.
	//
	// For insance, database or validation.
	// It can be used to identify or filter the error types in the logs.
	errType string

	// Original error that is returned by the function.
	//
	// For instance, the error returned by the database connection.
	// It can be used in the logs to help the developer.
	original error

	// Message to be returned in the error.
	//
	// This message is the curated message for the clients.
	// It can be used to display the error message to the clients.
	// For instance, the error message for database connection.
	// A database connection could be failed due to multiple reasons.
	// This message could be common for all of those reasons.
	message string

	// Status code to be set in the error.
	//
	// This can be used to set the status code in the http response.
	// For instance, the status code for database related errors could
	// be 500.
	status int

	// Attributes to be returned in the error.
	//
	// This can be used to set the attributes in the http response.
	// For instance, the attributes for validation errors could be
	// the list of fields that are invalid.
	attr map[string]interface{}
}

// New returns the error with the given error type and error.
func New(err error, errType string, opts ...Option) *Error {
	e := &Error{
		errType:  errType,
		original: err,
		message:  err.Error(),
	}

	e.merge(opts...)

	// Set the default status code if not set.
	if e.status == 0 {
		e.status = http.StatusInternalServerError
	}

	return e
}

// Error returns the error set in the message.
func (e *Error) Error() string {
	return e.message
}

// Type returns the type set in the error.
func (e *Error) Type() string {
	return e.errType
}

// OriginalError returns the original error.
func (e *Error) OriginalError() string {
	return e.original.Error()
}

// StatusCode returns the status code set in the error.
func (e *Error) StatusCode() int {
	return e.status
}

// Attributes returns the attributes set in the error.
func (e *Error) Attributes() map[string]interface{} {
	return e.attr
}
