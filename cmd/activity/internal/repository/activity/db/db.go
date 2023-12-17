package activitydb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iamsumit/sample-go-app/activity/internal/repository/activity"
	"github.com/iamsumit/sample-go-app/pkg/db"
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

// Create creates a new activity.
func (r *Repository) Create(ctx context.Context, activity activity.Activity) (*activity.Activity, error) {
	// To verify the database error.
	if activity.App.Name == "error" {
		return nil, db.ErrInternal(
			errors.New("app name is wrong"),
		)
	}

	return nil, nil
}
