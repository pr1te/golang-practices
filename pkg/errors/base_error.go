package errors

import (
	goerrors "github.com/go-errors/errors"
)

type Exception struct {
	Code    int           `json:"code" default:"500000"`
	Message string        `json:"message"`
	Details []interface{} `json:"details" default:"[]"`
}

func (err *Exception) Error() string {
	return err.Message
}

func New(message string, code int, details []interface{}) *goerrors.Error {
	msg := Messages[code]

	if len(message) > 0 {
		msg = message
	}

	err := []any{}

	if len(details) > 0 {
		err = details
	}

	excep := &Exception{
		Code:    code,
		Message: msg,
		Details: err,
	}

	return goerrors.New(excep)
}
