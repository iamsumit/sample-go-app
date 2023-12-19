package slogger

import (
	"log/slog"

	"github.com/jussi-kalliokoski/slogdriver"
	"github.com/lmittmann/tint"
)

// handler returns a new slog.Handler instance.
func handler(cfg cfg) slog.Handler {
	//--------------------------------------------------
	// Gcloud Handler
	//--------------------------------------------------
	if cfg.gcloud {
		return gcloudHandler(cfg)
	}

	//--------------------------------------------------
	// Slog Provided Handler
	//--------------------------------------------------

	// Set the handler options.
	opts := &slog.HandlerOptions{
		Level:       cfg.level,
		AddSource:   cfg.source,
		ReplaceAttr: nil,
	}

	switch cfg.format {
	case JSON:
		return slog.NewJSONHandler(
			cfg.writer,
			opts,
		)

	case TINT:
		return tint.NewHandler(
			cfg.writer,
			&tint.Options{
				Level:       cfg.level,
				AddSource:   cfg.source,
				ReplaceAttr: nil,
			},
		)

	case TEXT:
		fallthrough
	case DefaultFormat:
		fallthrough
	default:
		return slog.NewTextHandler(
			cfg.writer,
			opts,
		)
	}
}

// gcloudHandler returns a new slog.Handler instance.
//
// It will use the slogdriver package to set the handler to slog.
func gcloudHandler(cfg cfg) slog.Handler {
	// Set the handler options.
	opts := slogdriver.Config{
		ProjectID: cfg.projectID,
		Level:     cfg.level,
	}

	return slogdriver.NewHandler(
		cfg.writer,
		opts,
	)
}
