package api

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendJSONError(w http.ResponseWriter, e *InternalError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(e)
}

func SendJSONErrorV2(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json")
	err := NewInternalError(e.Error())
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}
