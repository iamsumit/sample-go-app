package user

import "errors"

var (
	// ErrInvalidID is returned when an invalid ID is passed to a method.
	ErrInvalidID = errors.New("invalid id")

	// ErrUserNotFound is returned when the user is not found.
	ErrUserNotFound = errors.New("user not found")

	// ErrDuplicateEmail occurs when a duplicate email is used to create/update a user.
	ErrDuplicateEmail = errors.New("email already exists")
)
