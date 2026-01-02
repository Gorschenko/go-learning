package errors

import "net/http"

type ErrorCode string

const (
	CodeBadRequest          ErrorCode = "CodeBadRequest"
	CodeNotFound            ErrorCode = "CodeNotFound"
	CodeRequestTimeout      ErrorCode = "CodeRequestTimeout"
	CodeInternalServerError ErrorCode = "InternalServerError"
)

var codeToStatus = map[ErrorCode]int{
	CodeBadRequest:     http.StatusBadRequest,
	CodeRequestTimeout: http.StatusRequestTimeout,
}
