package errors

import (
	goerrors "github.com/go-errors/errors"
)

func NewBadRequest(message string, details ...[]interface{}) *goerrors.Error {
	var err []interface{}

	if len(details) > 0 {
		err = details[0]
	}

	excep := New(message, BAD_REQUEST, err)

	return excep
}
