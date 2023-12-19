package slogger

import (
	"io"
	"log/slog"
	"os"
)

// Option to set the options for the logger.
type Option func(cfg *cfg)

// newCfg returns a new cfg instance.
func newCfg(opts ...Option) cfg {
	cfg := cfg{
		writer:    os.Stdout,
		format:    TEXT,
		level:     slog.LevelDebug,
		group:     false,
		source:    true,
		gcloud:    false,
		projectID: "",
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

// WithWriter sets the writer to write the logs to.
func WithWriter(w io.Writer) Option {
	return func(cfg *cfg) {
		cfg.writer = w
	}
}

// WithFormat sets the format of the logger.
func WithFormat(f Format) Option {
	return func(cfg *cfg) {
		cfg.format = f
	}
}

// WithOutSource sets the source information flag to false.
func WithOutSource() Option {
	return func(cfg *cfg) {
		cfg.source = false
	}
}

// WithLevel sets the level of the logger.
func WithLevel(l slog.Level) Option {
	return func(cfg *cfg) {
		cfg.level = l
	}
}

// WithGroup sets the level of the logger.
func WithGroup() Option {
	return func(cfg *cfg) {
		cfg.group = true
	}
}

// WithGcloudHandler sets the gcloud handler for the logger.
//
// It will use the slogdriver package to set the handler to slog.
// It will ignore the format options and use the its own format.
func WithGcloudHandler() Option {
	return func(cfg *cfg) {
		cfg.gcloud = true
	}
}

// WithProject sets the project id for the gcloud handler.
func WithProject(projectID string) Option {
	return func(cfg *cfg) {
		cfg.projectID = projectID
	}
}
