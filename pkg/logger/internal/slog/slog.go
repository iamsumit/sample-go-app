package slog // import "github.com/iamsumit/sample-go-app/pkg/logger/internal/slog"

import (
	"os"

	"golang.org/x/exp/slog"
)

// The SLog handler uses the slog package and implements the Logger interface.
type Handler struct {
	logger *slog.Logger
}

// Returns the slog logger object for the given log format.
func New(format string) *Handler {
	switch format {
	case "text":
		return getSLogTextLogger()
	case "json":
		return getSLogJSONLogger()
	default:
		return getSLogTextLogger()
	}
}

// Returns the slog logger object for the text log format.
func getSLogTextLogger() *Handler {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &Handler{logger}
}

// Returns the slog logger object for the JSON log format.
func getSLogJSONLogger() *Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Handler{logger}
}
