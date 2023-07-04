package exceptions

type Exception struct {
	Code    int           `json:"code" default:"500000"`
	Type    string        `json:"type"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors" default:"[]"`
}

func (err *Exception) Error() string {
	return err.Message
}

func New(message string, code int, errorType string, errors []interface{}) *Exception {
	msg := Messages[code]

	if len(msg) > 0 {
		msg = message
	}

	err := []any{}

	if len(errors) > 0 {
		err = errors
	}

	excep := &Exception{
		Code:    code,
		Type:    errorType,
		Message: msg,
		Errors:  err,
	}

	return excep
}
