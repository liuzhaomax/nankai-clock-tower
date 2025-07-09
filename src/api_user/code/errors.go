package code

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(code int, msg string) ApiError {
	return ApiError{
		Code:    code,
		Message: msg,
	}
}

func (code ApiError) Error() string {
	return code.Message
}

var (
	InternalErr       = Error(1000, "internal error")
	DBOperationFailed = Error(2000, "DB operation failed")
)
