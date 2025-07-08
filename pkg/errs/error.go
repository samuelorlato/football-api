package errs

import "net/http"

type Error struct {
	Message string `json:"erro"`
	Code    int    `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

func new(message string, code int) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func NewBadRequestError(message string) *Error {
	return new(message, http.StatusBadRequest)
}

func NewNotFoundError(message string) *Error {
	return new(message, http.StatusNotFound)
}

func NewUnprocessableContentError(message string) *Error {
	return new(message, http.StatusUnprocessableEntity)
}

func NewInternalServerError() *Error {
	return new("um erro inesperado aconteceu", http.StatusInternalServerError)
}
