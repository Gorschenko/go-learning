package api

import "net/http"

const (
	CodeBadRequest          string = "CodeBadRequest"
	CodeUnauthorized        string = "CodeUnauthorized"
	CodeForbidden           string = "CodeForbidden"
	CodeNotFound            string = "CodeNotFound"
	CodeAlreadyExists       string = "CodeAlreadyExists"
	CodeRequestTimeout      string = "CodeRequestTimeout"
	CodeInternalServerError string = "InternalServerError"
)

var codeToStatus = map[string]int{
	CodeBadRequest:          http.StatusBadRequest,
	CodeForbidden:           http.StatusForbidden,
	CodeNotFound:            http.StatusNotFound,
	CodeUnauthorized:        http.StatusUnauthorized,
	CodeAlreadyExists:       http.StatusConflict,
	CodeRequestTimeout:      http.StatusRequestTimeout,
	CodeInternalServerError: http.StatusInternalServerError,
}
