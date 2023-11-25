package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// The supported database type by this package.
type DBType int

const (
	UnknownDB DBType = iota
	MySQL     DBType = iota + 1
)

// Or use enum generator package.
func (t DBType) String() string {
	switch t {
	case MySQL:
		return "mysql"
	}

	return "unknown"
}

// The configuration required to initiate a database connection.
type Config struct {
	Type     DBType
	Name     string
	User     string
	Password string
	Port     int
	Host     string
}

// Initiate a database connection and retunrs the connection object.
func Handler(config *Config) (*sql.DB, error) {
	dsn := getDSN(config)

	db, err := sql.Open(config.Type.String(), dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

// Get the DSN string for the given database type.
func getDSN(config *Config) string {
	switch config.Type {
	case MySQL:
		return getMySQLDSN(config)
	}

	return ""
}

// Get the MySQL DSN string.
func getMySQLDSN(config *Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
}
