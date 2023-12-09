package store

import "errors"

var (
	// ErrDuplicateEmail occurs when a duplicate email is used to create/update a user.
	ErrorDuplicateEmail = errors.New("email already exists")
)
