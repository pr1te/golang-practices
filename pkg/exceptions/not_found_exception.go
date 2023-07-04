package exceptions

func NewNotFoundException(message string, errors ...[]interface{}) *Exception {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, NOT_FOUND, "NotFound", err)

	return excep
}
