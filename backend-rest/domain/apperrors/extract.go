package apperrors

import "log"

func Extract(err error) (int, string) {
	log.Println(err)
	customError, ok := err.(*Error)
	if !ok {
		return 500, "Internal Server Error"
	}
	return customError.Code, customError.Message
}
