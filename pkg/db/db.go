package db // import "github.com/iamsumit/sample-go-app/pkg/db"

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// The supported database type by this package.
type Type int

const (
	UnknownDB Type = iota
	MySQL     Type = iota + 1
)

// Or use enum generator package.
func (t Type) String() string {
	switch t {
	case MySQL:
		return "mysql"
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
	}

	return "", errors.New("Unsupported database type")
}

// Get the MySQL DSN string.
func getMySQLDSN(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
}
