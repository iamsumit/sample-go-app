package logger

import (
	"github.com/iamsumit/sample-go-app/pkg/logger/internal/slog"
)

// The supported logger type by this package.
type Type int

const (
	UnknownType Type = iota
	SLog        Type = iota + 1
)

// The supported log format for this package.
type Format int

const (
	UnknownFormat Format = iota
	Text          Format = iota + 1
	JSON          Format = iota + 2
)

// We can use enum code generator here as well.
func (lf Format) String() string {
	switch lf {
	case Text:
		return "text"
	case JSON:
		return "json"
	default:
		return "json"
	}
}

// The configuration required to initiate a logger object.
type Config struct {
	Type   Type
	Format Format
}

// Logger is the interface for every logger to use.
type Logger interface {
	Info(string, ...interface{})
	Warning(string, ...interface{})
	Error(string, ...interface{})
}

// Option type to use the options pattern to set the configuration.
type Option func(*Config)

// New returns a new logger object by using the given options.
func New(opts ...Option) (Logger, error) {
	config := newConfig(opts...)

	logger, err := getLogger(config)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// Get the logger object for the given logger type and format.
func getLogger(config *Config) (Logger, error) {
	switch config.Type {
	case SLog:
		return slog.New(config.Format.String())
	}

	return nil, nil
}
