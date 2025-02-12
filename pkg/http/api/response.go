package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Ok    bool   `json:"ok"`
	Body  any    `json:"body,omitempty"`
	Error *Error `json:"error,omitempty"`
}

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, body any, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	response := &Response{
		Ok:   true,
		Body: body,
	}
	json.NewEncoder(w).Encode(response)
}

func Err(w http.ResponseWriter, errMsg string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	response := &Response{
		Ok: false,
		Error: &Error{
			Status:  http.StatusText(statusCode),
			Message: errMsg,
		},
	}
	json.NewEncoder(w).Encode(response)
}
