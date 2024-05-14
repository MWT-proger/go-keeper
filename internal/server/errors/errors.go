package errors

import (
	"net/http"
)

func NewServicesError(text string, httpCode int) *ServicesError {
	return &ServicesError{text, httpCode}
}

type ServicesError struct {
	s        string
	HTTPCode int
}

func (e *ServicesError) Error() string {
	return e.s
}

var GetUserServicesError = NewServicesError(
	"не получилось обработать запрос получения пользователя",
	http.StatusInternalServerError,
)

var UserNotFoundServicesError = NewServicesError(
	"пользователь не авторизован",
	http.StatusUnauthorized,
)

var UserExistsServicesError = NewServicesError(
	"пользователь уже существует",
	http.StatusConflict,
)

var InternalServicesError = NewServicesError(
	"внутренняя ошибка сервера",
	http.StatusInternalServerError,
)
