package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iamsumit/sample-go-app/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// ExecContext is a wrapper to the sql.ExecContext.
//
// It accepts the context, sql.DB object, operation name, query and arguments.
// Operation name is used for logging and tracing purposes.
// It handles the error and tracing.
// It returns the sql.Result object and error.
func ExecContext(ctx context.Context, s *sql.DB, op string, query string, args ...any) (sql.Result, error) {
	// Start the tracing here.
	ctx, span := tracer.Global("db.ExecContext").Start(ctx, op)
	defer span.End()

	span.SetAttributes(
		attribute.String("db.query", query),
		attribute.String("db.args", fmt.Sprintf("%v", args)),
	)

	result, err := s.ExecContext(ctx, query, args...)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, ErrInternal(
			fmt.Errorf("%s error: %v", op, err),
		)
	}

	return result, nil
}

// QueryRowContext is a wrapper to the sql.QueryRowContext.
//
// It accepts the context, sql.DB object, operation name, query and arguments.
// Operation name is used for logging and tracing purposes.
// It handles the error and tracing.
// It returns the sql.Row object and error.
func QueryRowContext(ctx context.Context, s *sql.DB, op string, query string, args ...any) (*sql.Row, error) {
	// Start the tracing here.
	ctx, span := tracer.Global("db.QueryRowContext").Start(ctx, op)
	defer span.End()

	span.SetAttributes(
		attribute.String("db.query", query),
		attribute.String("db.args", fmt.Sprintf("%v", args)),
	)

	r := s.QueryRowContext(ctx, query, args...)

	err := r.Err()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, ErrInternal(
			fmt.Errorf("%s error: %v", op, err),
		)
	}

	return r, nil
}
