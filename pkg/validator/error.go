package validator

// FieldError is to keep track of validation error for particular field.
type FieldError struct {
	Field string
	Error string
}

// FieldErrors is a collection of FieldError.
type FieldErrors struct {
	Msg        string
	FieldError []FieldError
}

// Error returns a string representation of the error.
func (fe FieldErrors) Error() string {
	return fe.Msg
}

// FieldErrors returns a map of field name and error.
func (fe FieldErrors) FieldErrors() map[string]string {
	t := map[string]string{}
	for _, v := range fe.FieldError {
		t[v.Field] = v.Error
	}

	return t
}
