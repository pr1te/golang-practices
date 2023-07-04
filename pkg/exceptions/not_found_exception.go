package exceptions

import "github.com/go-errors/errors"

func NewNotFoundException(message string, errors ...[]interface{}) *errors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, NOT_FOUND, err)

	return excep
}
