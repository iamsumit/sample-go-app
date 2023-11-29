package slog

// Log the given message with the given attributes.
func (h *Handler) Info(message string, attr ...interface{}) {
	h.logger.Info(message, attr...)
}

// Log the given message with the given attributes.
func (h *Handler) Warning(message string, attr ...interface{}) {
	h.logger.Warn(message, attr...)
}

// Log the given message with the given attributes.
func (h *Handler) Error(message string, attr ...interface{}) {
	h.logger.Error(message, attr...)
}
