package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

// The SLog handler uses the slog package and implementes the Logger interface.
type SlogHandler struct {
	logger *slog.Logger
}

// Returns the slog logger object for the given log format.
func getSLogLogger(logFormatType LogFormatType) (*SlogHandler, error) {
	switch logFormatType {
	case Text:
		return getSLogTextLogger()
	case JSON:
		return getSLogJSONLogger()
	default:
		return getSLogTextLogger()
	}
}

// Returns the slog logger object for the text log format.
func getSLogTextLogger() (*SlogHandler, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &SlogHandler{logger}, nil
}

// Returns the slog logger object for the JSON log format.
func getSLogJSONLogger() (*SlogHandler, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &SlogHandler{logger}, nil
}

// Log the given message with the given attributes.
func (h *SlogHandler) Info(message string, attr ...interface{}) {
	h.logger.Info(message, attr...)
}

// Log the given message with the given attributes.
func (h *SlogHandler) Warning(message string, attr ...interface{}) {
	h.logger.Warn(message, attr...)
}

// Log the given message with the given attributes.
func (h *SlogHandler) Error(message string, attr ...interface{}) {
	h.logger.Error(message, attr...)
}
