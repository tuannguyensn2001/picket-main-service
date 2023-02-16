package app

import (
	"net/http"
	"strings"
)

type Error struct {
	Message    string `json:"message"`
	Code       int    `json:"code"`
	StatusCode int    `json:"status_code"`
}

func (e Error) Error() string {
	return e.Message
}

func NewRawError(message string, statusCode int) error {
	return &Error{
		Message:    message,
		Code:       1,
		StatusCode: statusCode,
	}
}

func NewInternalError(err error) error {
	return &Error{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func NewBadRequestError(err error, msg ...string) error {
	if len(msg) > 0 {
		return &Error{
			Message:    strings.Join(msg, "-"),
			StatusCode: http.StatusBadRequest,
		}
	}
	return &Error{
		Message:    err.Error(),
		StatusCode: http.StatusBadRequest,
	}
}

func NewForbiddenError(err error, msg ...string) error {
	if len(msg) > 0 {
		return &Error{
			Message:    strings.Join(msg, "-"),
			StatusCode: http.StatusForbidden,
		}
	}
	return &Error{
		Message:    err.Error(),
		StatusCode: http.StatusForbidden,
	}
}
