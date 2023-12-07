package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Handler struct {
	v *validator.Validate
	t ut.Translator
}

var (
	// A default validator instance.
	h = New()
)

// New returns a new validator instance.
func New() *Handler {

	validate := validator.New()

	// Create a translator for english so the error messages are
	// more human-readable than technical.
	translator, _ := ut.New(en.New(), en.New()).GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, translator)

	return &Handler{
		v: validate,
		t: translator,
	}
}

// Validate validates the given struct and returns an error if found.
func Validate(val interface{}) error {
	if err := h.v.Struct(val); err != nil {
		// If the error is a validator error, convert it.
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return fmt.Errorf("validation failed: %w", err)
		}

		fields := FieldErrors{
			Msg: "field validation failed",
		}

		for _, v := range errs {
			field := append(fields.FieldError, FieldError{
				Field: strings.ToLower(v.Field()),
				Error: strings.ToLower(v.Translate(h.t)),
			})

			fields.FieldError = append(fields.FieldError, field...)
		}

		return fields
	}

	return nil
}
