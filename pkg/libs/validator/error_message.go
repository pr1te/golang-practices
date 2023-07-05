package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	REQUIRED = "the '{field}' field is required."

	// email
	EMAIL = "the '{field}' field must be a valid e-mail."

	// password
	PASSWORD = "the '{field}' field must contain 8 or more characters with a mix of letters, numbers, symbols(@$#!%*?&), and at least one capital letter."
)

func getError(err validator.FieldError) ValidateErrorDetail {
	element := ValidateErrorDetail{
		Field:    err.Field(),
		Type:     err.Tag(),
		Expected: nil,
		Actual:   nil,
		Message:  err.Error(),
	}

	switch err.Tag() {
	case "required":
		element.Message = strings.Replace(REQUIRED, "{field}", err.Field(), 1)

	case "email":
		element.Actual = err.Value()
		element.Message = strings.Replace(EMAIL, "{field}", err.Field(), 1)

	case "password":
		element.Message = strings.Replace(PASSWORD, "{field}", err.Field(), 1)
	}

	return element
}
