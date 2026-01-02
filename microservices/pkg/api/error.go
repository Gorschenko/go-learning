package api

import "net/http"

type InternalError struct {
	Code    string `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewInternalError(code string) *InternalError {
	status, ok := codeToStatus[code]
	message := ""

	if !ok {
		status = http.StatusInternalServerError
		code = CodeInternalServerError
		message = code
	}

	return &InternalError{
		Code:    code,
		Status:  status,
		Message: message,
	}
}

// Immutable версия - возвращает новый объект
func (e *InternalError) WithMessage(m string) *InternalError {
	// Создаем копию, чтобы избежать побочных эффектов
	newErr := *e
	newErr.Message = m
	return &newErr
}
