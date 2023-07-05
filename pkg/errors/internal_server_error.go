package errors

import (
	goerrors "github.com/go-errors/errors"
)

func NewInternalServerError(message string, errors ...[]interface{}) *goerrors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, INTERNAL_SERVER_ERROR, err)

	return excep
}
