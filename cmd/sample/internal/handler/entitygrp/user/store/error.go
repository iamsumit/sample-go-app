// File: error.go declares the errors used by the user store.
// package store

package store

import (
	"errors"
	"net/http"

	"github.com/iamsumit/sample-go-app/pkg/db"
)

var (
	// ErrDuplicateEmail occurs when a duplicate email is used to create/update a user.
	ErrDuplicateEmail = db.NewError(
		errors.New("email already exists"),
		http.StatusConflict,
	)

	// ErrUserNotFound is returned when the user is not found.
	ErrUserNotFound = db.NewError(
		errors.New("user not found"),
		http.StatusNotFound,
	)
)
