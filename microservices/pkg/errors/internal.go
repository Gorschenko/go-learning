package errors

import "net/http"

type InternalError struct {
	Code    string `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewInternalError(code string) *InternalError {
	status, ok := codeToStatus[code]

	if !ok {
		status = http.StatusInternalServerError
	}

	return &InternalError{
		Code:   code,
		Status: status,
	}
}

func (e *InternalError) WithMessage(m string) *InternalError {
	e.Message = m
	return e
}
