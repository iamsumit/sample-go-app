package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/lib/pq"              // Postgres driver
)

// The supported database type by this package.
type Type int

const (
	UnknownDB Type = iota
	MySQL     Type = iota + 1
	Postgres  Type = iota + 2
)

// Or use enum generator package.
func (t Type) String() string {
	switch t {
	case MySQL:
		return "mysql"
	case Postgres:
		return "postgres"
	}

	return "unknown"
}

// The configuration required to initiate a database connection.
type Config struct {
	Type     Type
	Name     string
	User     string
	Password string
	Port     int
	Host     string
}

// Initiate a database connection and retunrs the connection object.
func New(config Config) (*sql.DB, error) {
	dsn, err := getDSN(config)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(config.Type.String(), dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

// Get the DSN string for the given database type.
func getDSN(config Config) (string, error) {
	switch config.Type {
	case MySQL:
		return getMySQLDSN(config), nil
	case Postgres:
		return getPGSQLDSN(config), nil
	}

	return "", errors.New("Unsupported database type")
}

// Get the MySQL DSN string.
func getMySQLDSN(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Name)
}

// Get the PGSQL DSN string.
func getPGSQLDSN(config Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Name)
}
