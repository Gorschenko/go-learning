package api

import "net/http"

const (
	CodeBadRequest          string = "CodeBadRequest"
	CodeNotFound            string = "CodeNotFound"
	CodeRequestTimeout      string = "CodeRequestTimeout"
	CodeInternalServerError string = "InternalServerError"
)

var codeToStatus = map[string]int{
	CodeBadRequest:          http.StatusBadRequest,
	CodeNotFound:            http.StatusNotFound,
	CodeRequestTimeout:      http.StatusRequestTimeout,
	CodeInternalServerError: http.StatusInternalServerError,
}
