package exceptions

import "github.com/go-errors/errors"

func NewValidationErrorException(message string, errors ...[]interface{}) *errors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, VALIDATION_ERROR, err)

	return excep
}
