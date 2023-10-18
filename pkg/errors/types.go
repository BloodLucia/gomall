package errors

import "net/http"

func BadRequest(msg string) *Error {
	return New(http.StatusBadRequest, msg)
}

func Unauthorized() *Error {
	return New(http.StatusUnauthorized, "Unauthorized")
}

func Forbidden() *Error {
	return New(http.StatusForbidden, "Forbidden")
}

func NotFound(msg string) *Error {
	return New(http.StatusNotFound, msg)
}

func InternalServer() *Error {
	return New(http.StatusInternalServerError, "服务器内部错误")
}

func IsInternalServer(err *Error) bool {
	return err.Code == http.StatusInternalServerError
}
