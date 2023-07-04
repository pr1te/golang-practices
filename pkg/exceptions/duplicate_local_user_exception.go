package exceptions

import "github.com/go-errors/errors"

func NewDuplicateLocalUserException(message string, errors ...[]interface{}) *errors.Error {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	return New(message, DUPLICATE_LOCAL_USER, err)
}
