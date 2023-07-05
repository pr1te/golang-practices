package errors

import (
	goerrors "github.com/go-errors/errors"
)

func newValidationFailed(message string, errors ...[]interface{}) *goerrors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, VALIDATION_ERROR, err)

	return excep
}
