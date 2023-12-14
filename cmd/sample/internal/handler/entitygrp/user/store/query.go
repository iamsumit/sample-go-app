package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/iamsumit/sample-go-app/pkg/db"
)

// ByID returns the user for the given id.
//
//nolint:dupl
func (h *Handler) ByID(ctx context.Context, id int) (*User, error) {
	query, args, err := squirrel.
		Select(
			"users.id",
			"users.name",
			"users.email",
			"users.is_active",
			"user_settings.is_subscribed",
			"user_settings.biography",
			"user_settings.date_of_birth",
			"users.created_at",
			"users.updated_at",
		).
		From("users").
		LeftJoin("user_settings ON users.id = user_settings.uid").
		Where(
			squirrel.And{
				squirrel.Eq{"users.id": id},
				squirrel.Eq{"users.deleted_at": nil},
			},
		).ToSql()
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("user.ByID - unable to build select query: %w", err),
		)
	}

	user := new(User)
	row, err := db.QueryRowContext(ctx, h.db, "user.ByID", query, args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.IsActive,
		&user.Settings.IsSubscribed,
		&user.Settings.Biography,
		&user.Settings.DateOfBirth,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, db.ErrInternal(
			fmt.Errorf("user.ByID - unable to query data: %w", err),
		)
	}

	return user, nil
}

// ByEmail returns the user for the given email.
//
//nolint:dupl
func (h *Handler) ByEmail(ctx context.Context, email string) (*User, error) {
	query, args, err := squirrel.
		Select(
			"users.id",
			"users.name",
			"users.email",
			"users.is_active",
			"user_settings.is_subscribed",
			"user_settings.biography",
			"user_settings.date_of_birth",
			"users.created_at",
			"users.updated_at",
		).
		From("users").
		Join("user_settings ON users.id = user_settings.uid").
		Where(
			squirrel.And{
				squirrel.Eq{"users.email": email},
				squirrel.Eq{"users.deleted_at": nil},
			},
		).ToSql()
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("user.ByEmail - unable to build select query: %w", err),
		)
	}

	user := new(User)

	row, err := db.QueryRowContext(ctx, h.db, "user.ByEmail", query, args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.IsActive,
		&user.Settings.IsSubscribed,
		&user.Settings.Biography,
		&user.Settings.DateOfBirth,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, db.ErrInternal(
			fmt.Errorf("user.ByEmail - unable to query data: %w", err),
		)
	}

	return user, nil
}

// Create adds a new user in the database after validating the input.
func (h *Handler) Create(ctx context.Context, user User) (*User, error) {
	if user.Email != nil {
		// If email is provided, it must be unique.
		existingUser, err := h.ByEmail(ctx, *user.Email)
		if err != nil && !errors.Is(err, ErrUserNotFound) {

			return nil, err
		}

		if existingUser != nil {

			return nil, ErrDuplicateEmail
		}
	}

	// Build the insert query for users table.
	query, args, err := squirrel.Insert("users").
		Columns("name", "email", "is_active").
		Values(
			user.Name,
			user.Email,
			user.IsActive,
		).ToSql()
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("user.Create - unable to build user insert query: %w", err),
		)
	}

	// Insert the user in the database.
	result, err := db.ExecContext(ctx, h.db, "user.Create", query, args...)
	if err != nil {
		return nil, err
	}

	// Get the last insert id.
	id, err := result.LastInsertId()
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("user.Create: unable to get last insert id: %w", err),
		)
	}

	// Build the insert query for user_settings table.
	query, args, err = squirrel.Insert("user_settings").
		Columns("uid", "is_subscribed", "biography", "date_of_birth").
		Values(
			int(id),
			user.Settings.IsSubscribed,
			user.Settings.Biography,
			user.Settings.DateOfBirth,
		).ToSql()
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("user.Create - unable to build user settings insert query: %w", err),
		)
	}

	// Insert the user settings in the database.
	_, err = db.QueryRowContext(ctx, h.db, "user.Create", query, args...)
	if err != nil {
		return nil, err
	}

	// Get the last created user.
	createdUser, err := h.ByID(ctx, int(id))
	if err != nil {

		return nil, err
	}

	return createdUser, nil
}
