package validator

type Error error

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
func (fe FieldErrors) FieldErrors() map[string]interface{} {
	t := map[string]interface{}{}
	for _, v := range fe.FieldError {
		t[v.Field] = v.Error
	}

	return t
}

// Attributes returns a map of field name and error.
func Attributes(e Error) map[string]interface{} {
	switch e.(type) {
	case FieldErrors:
		return e.(FieldErrors).FieldErrors()
	}

	return nil
}
