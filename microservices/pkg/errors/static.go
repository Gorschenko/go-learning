package errors

import "net/http"

type ErrorCode string

const (
	CodeBadRequest     ErrorCode = "CodeBadRequest"
	CodeRequestTimeout ErrorCode = "CodeRequestTimeout"
)

var codeToStatus = map[ErrorCode]int{
	CodeBadRequest:     http.StatusBadRequest,
	CodeRequestTimeout: http.StatusRequestTimeout,
}
