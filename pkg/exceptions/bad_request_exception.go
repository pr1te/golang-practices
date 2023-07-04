package exceptions

func NewBadRequestException(message string, errors ...[]interface{}) *Exception {
	var err []interface{}

	if len(errors) > 0 {
		err = errors[0]
	}

	excep := New(message, BAD_REQUEST, "BadRequest", err)

	return excep
}
