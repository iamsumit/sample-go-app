package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// BirthDateValidator is a custom validator function for validating the date in "YYYY-MM-DD" format.
func BirthDateValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
