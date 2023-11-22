package logger

// The supported logger type by this package.
type LoggerType int

const (
	Unknown LoggerType = iota
	SLog    LoggerType = iota + 1
)

// The supported log format for this package.
type LogFormatType int

const (
	Default LogFormatType = iota
	Text    LogFormatType = iota + 1
	JSON    LogFormatType = iota + 1
)

// The configuration required to initiate a logger object.
type Config struct {
	LoggerType LoggerType
	LogFormat  LogFormatType
}

// The logger interface implemented by all the loggers.
type Logger interface {
	Info(string, ...interface{})
	Warning(string, ...interface{})
	Error(string, ...interface{})
}

// Get the logger object for the given logger type and format.
func getLogger(config *Config) (Logger, error) {
	switch config.LoggerType {
	case SLog:
		return getSLogLogger(config.LogFormat)
	}

	return nil, nil
}

// Initiate a logger object and returns it.
func New(config *Config) (Logger, error) {
	logger, err := getLogger(config)
	if err != nil {
		return nil, err
	}

	return logger, nil
}
