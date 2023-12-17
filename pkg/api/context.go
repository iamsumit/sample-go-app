package api

import (
	"context"
	"errors"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// key is how request values are stored/retrieved.
const key ctxKey = 1

// ContextValues represent state for each request.
type ContextValues struct {
	StatusCode int
	IsError    bool
}

// GetContextValues returns the values from the context.
func GetContextValues(ctx context.Context) (*ContextValues, error) {
	v, ok := ctx.Value(key).(*ContextValues)
	if !ok {
		return nil, errors.New("api value missing from context")
	}
	return v, nil
}

// SetStatusCode sets the status code back into the context.
func SetStatusCode(ctx context.Context, statusCode int) error {
	v, ok := ctx.Value(key).(*ContextValues)
	if !ok {
		return errors.New("status code value missing from context")
	}
	v.StatusCode = statusCode
	return nil
}

// SetIsError sets the error code back into the context.
func SetIsError(ctx context.Context) error {
	v, ok := ctx.Value(key).(*ContextValues)
	if !ok {
		return errors.New("isError value missing from context")
	}
	v.IsError = true
	return nil
}
