package errors

import (
	"fmt"
	"net/http"
)

// Error 自定义错误类型
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error: code=%d, message=%s, err=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("error: code=%d, message=%s", e.Code, e.Message)
}

// New 创建新的错误
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装已有错误
func Wrap(err error, code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// 预定义错误码
const (
	// 系统级错误码 (1000-1999)
	ErrInternalServer = 1000
	ErrInvalidParam   = 1001
	ErrUnauthorized   = 1002
	ErrForbidden      = 1003
	ErrNotFound       = 1004

	// 业务级错误码 (2000-2999)
	ErrUserNotFound     = 2000
	ErrUserAlreadyExist = 2001
	ErrInvalidPassword  = 2002
)

// 预定义错误
var (
	InternalServerError = New(ErrInternalServer, "Internal server error")
	InvalidParamError   = New(ErrInvalidParam, "Invalid parameter")
	UnauthorizedError   = New(ErrUnauthorized, "Unauthorized")
	ForbiddenError      = New(ErrForbidden, "Forbidden")
	NotFoundError       = New(ErrNotFound, "Not found")

	UserNotFoundError     = New(ErrUserNotFound, "User not found")
	UserAlreadyExistError = New(ErrUserAlreadyExist, "User already exists")
	InvalidPasswordError  = New(ErrInvalidPassword, "Invalid password")
)

// HTTPStatus 获取错误对应的 HTTP 状态码
func (e *Error) HTTPStatus() int {
	switch e.Code {
	case ErrInternalServer:
		return http.StatusInternalServerError
	case ErrInvalidParam:
		return http.StatusBadRequest
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrForbidden:
		return http.StatusForbidden
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
} 