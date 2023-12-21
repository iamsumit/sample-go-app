// Package logger for logger functions
package slogger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"runtime/debug"
	"time"

	"github.com/lmittmann/tint"
)

// Handler is the logger instance.
type Handler struct {
	// writer is the writer to write the logs to.
	// It will set the writer to the provided value.
	//
	// If not provided, it will default to os.Stdout.
	writer io.Writer

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
}

// New returns a new Handler instance.
//
// It can be used to fetch different kinds of logger.
func New(w io.Writer, source bool) *Handler {
	return &Handler{
		writer: w,
		level:  slog.LevelDebug,
		source: source,
	}
}

// WithFormat returns a new slog.Logger instance with the provided format.
//
// It will return a logger with text format if the format is not supported.
//
// format: The format of the logger.
// Supported formats:
//   - json
//   - tint
//   - text
func WithFormat(w io.Writer, format string) *slog.Logger {
	var log *slog.Logger

	lHandler := New(w, true)

	switch format {
	case "json":
		log = lHandler.JSONLogger()

	case "tint":
		log = lHandler.TintLogger()

	case "text":
		fallthrough
	default:
		log = lHandler.TextLogger()
	}

	return WithGroup(log)
}

// WithGroup returns a new slog.Logger instance with the group information.
//
// It will add the following information:
// - pid: the process id of the application.
// - go_version: the version of the Go runtime.
func WithGroup(l *slog.Logger) *slog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()

	return l.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
}

// TextLogger returns a new slog.Logger instance with text handler.
func (h *Handler) TextLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:       h.level,
		AddSource:   h.source,
		ReplaceAttr: nil,
	}

	return slog.New(
		slog.NewTextHandler(
			h.writer,
			opts,
		),
	)
}

// JSONLogger returns a new slog.Logger instance with json handler.
func (h *Handler) JSONLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:       h.level,
		AddSource:   h.source,
		ReplaceAttr: nil,
	}

	return slog.New(
		slog.NewJSONHandler(
			h.writer,
			opts,
		),
	)
}

// TintLogger returns a new slog.Logger instance with tint handler.
func (h *Handler) TintLogger() *slog.Logger {
	opts := &tint.Options{
		Level:       h.level,
		AddSource:   h.source,
		ReplaceAttr: nil,
		TimeFormat:  time.StampMilli,
		NoColor:     false,
	}

	return slog.New(
		tint.NewHandler(
			h.writer,
			opts,
		),
	)
}

// ServerLogger returns a new log.Logger instance with level error.
//
// It can be used with http.Server so that it can log error in same format as
// our logger.
func ServerLogger(l *slog.Logger) *log.Logger {
	return slog.NewLogLogger(l.Handler(), slog.LevelError)
}
