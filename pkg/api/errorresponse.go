package api

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	// in:body
	//
	//example: data is not in proper format
	Error string `json:"error"`
	// in:body
	//
	//example: {"field": "error message for this specific field"}
	Data map[string]interface{} `json:"data,omitempty"`
}

// RequestError is used to pass an error during the request through the
// application with web specific context.
type RequestError struct {
	Err    error
	Status int
}

// NewRequestError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewRequestError(err error, status int) error {
	return &RequestError{
		Err:    err,
		Status: status,
	}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (re *RequestError) Error() string {
	return re.Err.Error()
}
