package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	v *validator.Validate
}

type ValidateErrorDetail struct {
	FailedField string
	Tag         string
	Value       string
}

func (v *Validator) ValidateStruct(stru interface{}) []ValidateErrorDetail {
	var errors []ValidateErrorDetail
	err := v.v.Struct(stru)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ValidateErrorDetail{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}

			errors = append(errors, element)
		}
	}

	return errors
}

func New() *Validator {
	v := validator.New()

	return &Validator{v}
}
