package activitydb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	activitypkg "github.com/iamsumit/sample-go-app/activity/internal/repository/activity"
	"github.com/iamsumit/sample-go-app/pkg/db"
	"github.com/iamsumit/sample-go-app/pkg/util/strings"
)

// Get app by alias.
func (r *Repository) AppByAlias(ctx context.Context, alias string) (*activitypkg.App, error) {
	// -------------------------------------------------------------------
	// App by alias
	// -------------------------------------------------------------------
	query, args, err := squirrel.
		Select(
			"app.id",
			"app.name",
			"app.alias",
			"app.created_at",
			"app.updated_at",
		).
		From("app").
		Where(
			squirrel.And{
				squirrel.Eq{"app.alias": alias},
				squirrel.Eq{"app.deleted_at": nil},
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("activitydb.AppByAlias - unable to build select query: %w", err),
		)
	}

	app := new(activitypkg.App)

	row, err := db.QueryRowContext(ctx, r.db, "activity.AppByAlias", query, args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&app.ID,
		&app.Name,
		&app.Alias,
		&app.CreatedAt,
		&app.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, activitypkg.ErrAppNotFound
		}

		return nil, db.ErrInternal(
			fmt.Errorf("activity.AppByAlias - unable to scan data: %w", err),
		)
	}

	return app, nil
}

// CreateApp creates a new app.
func (r *Repository) CreateApp(ctx context.Context, app activitypkg.App) (*activitypkg.App, error) {
	// -------------------------------------------------------------------
	// Create app
	// -------------------------------------------------------------------
	query, args, err := squirrel.
		Insert("app").
		Columns(
			"name",
			"alias",
		).
		Values(
			app.Name,
			strings.ToAlias(app.Name),
		).
		Suffix(
			"RETURNING \"id\", \"alias\", \"created_at\", \"updated_at\"",
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("activitydb.CreateApp - unable to build insert query: %w", err),
		)
	}

	row, err := db.QueryRowContext(ctx, r.db, "activity.CreateApp", query, args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&app.ID,
		&app.Alias,
		&app.CreatedAt,
		&app.UpdatedAt,
	)
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("activity.CreateApp - unable to scan data: %w", err),
		)
	}

	return &app, nil
}

// Create creates a new activity.
func (r *Repository) Create(ctx context.Context, activity activitypkg.Activity) (*activitypkg.Activity, error) {
	// -------------------------------------------------------------------
	// App by alias
	// -------------------------------------------------------------------

	appNameAlias := strings.ToAlias(activity.App.Name)

	app, err := r.AppByAlias(ctx, appNameAlias)
	if err != nil && !errors.Is(err, activitypkg.ErrAppNotFound) {
		return nil, err
	}

	if err != nil && errors.Is(err, activitypkg.ErrAppNotFound) {
		app, err = r.CreateApp(ctx, activity.App)
		if err != nil {
			return nil, err
		}
	}

	// -------------------------------------------------------------------
	// Create activity
	// -------------------------------------------------------------------

	query, args, err := squirrel.
		Insert("activity").
		Columns(
			"aid",
			"entity",
			"operation",
		).
		Values(
			app.ID,
			activity.Entity,
			activity.Operation,
		).
		Suffix(
			"RETURNING \"id\", \"created_at\", \"updated_at\"",
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("activitydb.Create - unable to build insert query: %w", err),
		)
	}

	row, err := db.QueryRowContext(ctx, r.db, "activity.Create", query, args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&activity.ID,
		&activity.CreatedAt,
		&activity.UpdatedAt,
	)
	if err != nil {
		return nil, db.ErrInternal(
			fmt.Errorf("activity.Create - unable to scan data: %w", err),
		)
	}

	activity.App = *app

	return &activity, nil
}
