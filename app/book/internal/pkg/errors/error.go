package errors

import "fmt"

type AppError struct {
	Code    int32  // 错误码，跟 grpc-status 一致，并且在HTTP中可映射成 http-status
	Reason  string // 错误原因，业务判定错误码
	Message string // 用户可读信息，可作为用户提示内容
}

func (e *AppError) Error() string {
	return e.Reason
}

func New(code int32, reason string, msg ...interface{}) *AppError {
	return &AppError{
		Code:    code,
		Reason:  reason,
		Message: format(msg...),
	}
}

func format(msg ...interface{}) string {
	if len(msg) > 0 {
		if str, ok := msg[0].(string); ok {
			return fmt.Sprintf(str, msg[1:])
		}
	}
	return ""
}
