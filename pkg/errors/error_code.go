package errors

const (
	// 500xxx
	INTERNAL_SERVER_ERROR = 500000

	// 400xxx
	BAD_REQUEST          = 400000
	DUPLICATE_LOCAL_USER = 400001
	VALIDATION_ERROR     = 400002

	// 401xxx
	UNAUTHORIZED = 401000

	// 404xxx
	NOT_FOUND = 404000
)
