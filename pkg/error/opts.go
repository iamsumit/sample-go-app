package error

// Merge the configurations by the config set in the options.
func (c *Error) merge(opts ...Option) {
	if c == nil {
		return
	}

	for _, f := range opts {
		f(c)
	}
}

// WithMessage sets the message in the error.
func WithMessage(message string) Option {
	return func(e *Error) {
		e.message = message
	}
}

// WithStatus sets the status in the error.
func WithStatus(status int) Option {
	return func(e *Error) {
		e.status = status
	}
}

// WithAttributes sets the attributes in the error.
func WithAttributes(attr map[string]interface{}) Option {
	return func(e *Error) {
		e.attr = attr
	}
}
