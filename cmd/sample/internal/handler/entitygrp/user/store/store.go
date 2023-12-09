// Package store provides the database support for the user entity.
//
// It can be used to retrieve and store user information by given methods.
// It provides User model to be used to pass and retrieve user information.
package store

import "database/sql"

// Handler holds the dependencies for the user store.
type Handler struct {
	db *sql.DB
}

// New returns a new user store.
func New(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}
