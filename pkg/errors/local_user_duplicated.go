package errors

import (
	goerrors "github.com/go-errors/errors"
)

func NewLocalUserDuplicated(message string, errors ...[]interface{}) *goerrors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	return New(message, DUPLICATE_LOCAL_USER, err)
}
