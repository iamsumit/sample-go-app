package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/iamsumit/sample-go-app/pkg/db"
)

// New opens a database connection with the given config.
func New(t *testing.T, cfg db.Config, testDBName string) (*sql.DB, func(), error) {
	// Create a new database with the given name.
	sqlDB, err := db.New(cfg)
	if err != nil {
		t.Fatalf("failed to initiate database instance: %v", err)
		return nil, nil, fmt.Errorf("failed to initiate database instance: %w", err)
	}

	t.Log("Database instance is ready!")

	// Ping the database to ensure it's working.
	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("failed to ping the database: %v", err)
		return nil, nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	t.Log("Database pinged successfully!")

	// Make sure we have sufficient permission for the db user.
	if _, err := sqlDB.ExecContext(context.Background(), "CREATE DATABASE IF NOT EXISTS "+testDBName); err != nil {
		t.Fatalf("creating database %s: %v", testDBName, err)
	}

	t.Log("database ready! closing current connection...")

	_ = sqlDB.Close()

	t.Log("database connection closed! opening new connection with test database ...")

	// connect to the newly created database.
	cfg.Name = testDBName
	sqlDB, err = db.New(cfg)
	if err != nil {
		t.Fatalf("failed to open the database %s: %v", testDBName, err)
		return nil, nil, fmt.Errorf("failed to open the database %s: %w", testDBName, err)
	}

	t.Log("Database setup is ready for testing!")

	teardown := func() {
		t.Helper()
		_ = sqlDB.Close()
	}

	return sqlDB, teardown, nil
}
