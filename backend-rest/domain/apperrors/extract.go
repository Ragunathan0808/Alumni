package apperrors

func Extract(err error) (int, string) {
	customError, ok := err.(*Error)
	if !ok {
		return 500, "Internal Server Error"
	}
	return customError.Code, customError.Message
}
