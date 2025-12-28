package errors

import "net/http"

const (
	CodeBadRequest     = "CodeBadRequest"
	CodeRequestTimeout = "CodeRequestTimeout"
)

var codeToStatus = map[string]int{
	CodeBadRequest:     http.StatusBadRequest,
	CodeRequestTimeout: http.StatusRequestTimeout,
}
