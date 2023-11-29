package logger

// Returns the default configuration for the logger.
func newConfig(opts ...Option) *Config {
	config := &Config{
		Type:   SLog,
		Format: JSON,
	}

	config.merge(opts...)

	return config
}

// Merge the configurations by the config set in the options.
func (c *Config) merge(opts ...Option) {
	if c == nil {
		return
	}

	for _, f := range opts {
		f(c)
	}
}

// WithSlogger sets the logger type to slog.
func WithSlogger() Option {
	return func(config *Config) {
		config.Type = SLog
	}
}

// WithTextFormat sets the log format to text.
func WithTextFormat() Option {
	return func(config *Config) {
		config.Format = Text
	}
}

// WithJSONFormat sets the log format to json.
func WithJSONFormat() Option {
	return func(config *Config) {
		config.Format = JSON
	}
}
