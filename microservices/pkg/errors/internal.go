package errors

import "net/http"

type InternalError struct {
	Code    ErrorCode `json:"code"`
	Status  int       `json:"status"`
	Message string    `json:"message"`
}

func NewInternalError(code ErrorCode) *InternalError {
	status, ok := codeToStatus[code]

	if !ok {
		status = http.StatusInternalServerError
	}

	return &InternalError{
		Code:   code,
		Status: status,
	}
}

// Реализация интерфейса error
func (e *InternalError) Error() string {
	return string(e.Code)
}

// Immutable версия - возвращает новый объект
func (e *InternalError) WithMessage(m string) *InternalError {
	// Создаем копию, чтобы избежать побочных эффектов
	newErr := *e
	newErr.Message = m
	return &newErr
}
