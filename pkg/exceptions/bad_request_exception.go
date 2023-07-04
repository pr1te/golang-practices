package exceptions

import "github.com/go-errors/errors"

func NewBadRequestException(message string, errors ...[]interface{}) *errors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, BAD_REQUEST, "BadRequest", err)

	return excep
}
