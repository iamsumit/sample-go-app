package activitydb

import (
	"database/sql"
)

// Repository holds the database repository for activity.
type Repository struct {
	db *sql.DB
}

// New returns a new instance of the activity repository.
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
