package validator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	errpkg "github.com/iamsumit/sample-go-app/pkg/error"
)

type Handler struct {
	v *validator.Validate
	t ut.Translator
}

// A default validator instance.
var h = New()

// New returns a new validator instance.
func New() *Handler {
	validate := validator.New()

	// Create a translator for english so the error messages are
	// more human-readable than technical.
	translator, _ := ut.New(en.New(), en.New()).GetTranslator("en")
	_ = entranslations.RegisterDefaultTranslations(validate, translator)

	// ---------------------------------------------------------
	// Birth date validators
	// ---------------------------------------------------------
	validate.RegisterValidation("birthDate", BirthDateValidator)
	validate.RegisterTranslation("birthDate", translator, func(ut ut.Translator) error {
		return ut.Add("birthDate", "{0} is not a valid format", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("birthDate", fe.Field())
		return t
	})

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
			return NewError(
				fmt.Errorf("validation failed: %w", err),
				http.StatusInternalServerError,
				nil,
			)
		}

		// Convert the errors to a map of field name and error.
		var attr = map[string]interface{}{}
		for k, v := range errs.Translate(h.t) {
			attr[k] = strings.ToLower(v)
		}

		return NewError(
			err,
			http.StatusBadRequest,
			attr,
			errpkg.WithMessage("field validation failed"),
		)
	}

	return nil
}
