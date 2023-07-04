package exceptions

import "github.com/go-errors/errors"

func NewInternalServerErrorException(message string, errors ...[]interface{}) *errors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, INTERNAL_SERVER_ERROR, "InternalServerError", err)

	return excep
}
