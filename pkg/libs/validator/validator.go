package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v *validator.Validate
}

type ValidateErrorDetail struct {
	Field    string `json:"field,omitempty"`
	Type     string `json:"type,omitempty"`
	Expected any    `json:"expected,omitempty"`
	Actual   any    `json:"actual,omitempty"`
	Message  string `json:"message,omitempty"`
}

func (v *Validator) ValidateStruct(stru interface{}) []ValidateErrorDetail {
	var errors []ValidateErrorDetail
	err := v.v.Struct(stru)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, getError(err))
		}
	}

	return errors
}

func New() *Validator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")

		if name == "-" {
			return ""
		}

		return name
	})

	// register custom validator
	v.RegisterValidation("password", validatePassword)

	return &Validator{v}
}
