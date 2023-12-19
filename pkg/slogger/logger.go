// Package logger for logger functions
package slogger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"runtime/debug"
)

// Format is the format of the logger.
type Format int

// Format constants.
const (
	// DefaultFormat is the default format of the logger.
	DefaultFormat Format = iota

	// Text is the text format of the logger.
	TEXT Format = iota + 1

	// JSON is the JSON format of the logger.
	JSON Format = iota + 2

	// TINT provides colored output to the text format.
	TINT Format = iota + 3
)

// cfg holds the configuration for slog.
type cfg struct {
	// writer is the writer to write the logs to.
	// It will set the writer to the provided value.
	//
	// If not provided, it will default to os.Stdout.
	writer io.Writer

	// format is the format of the logger.
	// It will be used to set the handler to slog.
	//
	// If not provided, it will default to TEXT.
	format Format

	// level is the level of the logger.
	// It will set Level to the provided value.
	//
	// If not provided, it will default to slog.LevelDebug.
	level slog.Level

	// source is the source information flag.
	// It will set AddSource to true.
	//
	// If not provided, it will default to true.
	source bool

	// group is the group information flag.
	// It will return the logger with basic grouping information.
	//
	// If not provided, it will default to false.
	group bool

	// gcloud is the gcloud handler flag.
	// It will use the slogdriver package to set the handler to slog.
	//
	// If not provided, it will default to false.
	gcloud bool

	// projectID is the project id for the gcloud handler.
	//
	// If not provided, it will use the empty string.
	projectID string
}

// New returns a new slog.Logger instance.
func New(opts ...Option) *slog.Logger {
	cfg := newCfg(opts...)

	log := slog.New(
		handler(cfg),
	)

	if cfg.group {
		return withGroup(log)
	}

	return log
}

// withGroup returns a new slog.Logger instance with the group information.
//
// It will add the following information:
// - pid: the process id of the application.
// - go_version: the version of the Go runtime.
func withGroup(l *slog.Logger) *slog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()

	return l.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
}

// ErrorLogger returns a new log.Logger instance with level error.
//
// It can be used with http.Server so that it can log error in same format as
// our logger.
func ErrorLogger(l *slog.Logger) *log.Logger {
	return slog.NewLogLogger(l.Handler(), slog.LevelError)
}
