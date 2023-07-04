package exceptions

import "github.com/go-errors/errors"

type Exception struct {
	Code    int           `json:"code" default:"500000"`
	Type    string        `json:"type"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors" default:"[]"`
}

func (err *Exception) Error() string {
	return err.Message
}

func New(message string, code int, errorType string, errs []interface{}) *errors.Error {
	msg := Messages[code]

	if len(msg) > 0 {
		msg = message
	}

	err := []any{}

	if len(errs) > 0 {
		err = errs
	}

	excep := &Exception{
		Code:    code,
		Type:    errorType,
		Message: msg,
		Errors:  err,
	}

	return errors.New(excep)
}
