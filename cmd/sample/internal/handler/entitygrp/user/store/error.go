// File: error.go declares the errors used by the user store.
// package store

package store

import "errors"

var (
	// ErrDuplicateEmail occurs when a duplicate email is used to create/update a user.
	ErrDuplicateEmail = errors.New("email already exists")

	// ErrUserNotFound is returned when the user is not found.
	ErrUserNotFound = errors.New("user not found")
)
