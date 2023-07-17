package errors

import (
	goerrors "github.com/go-errors/errors"
)

func NewLocalUserDuplicated(message string, details ...[]interface{}) *goerrors.Error {
	var err []interface{}

	if len(details) > 0 {
		err = details[0]
	}

	return New(message, DUPLICATE_LOCAL_USER, err)
}
