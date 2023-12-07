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
