package user

import "errors"

var (
	// ErrInvalidID is returned when an invalid ID is passed to a method.
	ErrInvalidID = errors.New("invalid id")

	// ErrUserNotFound is returned when the user is not found.
	ErrUserNotFound = errors.New("user not found")

	// ErrPayloadDecode is returned when the payload is not able to decode.
	ErrPayloadDecode = errors.New("unable to decode payload")
)
