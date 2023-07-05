package validator

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

func validatePassword(fl validator.FieldLevel) bool {
	regex := regexp2.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%#*?&])[A-Za-z\d@$#!%*?&]{8,}$`, regexp2.RE2)

	if matched, err := regex.MatchString(fl.Field().String()); matched && err == nil {
		return true
	}

	return false
}
