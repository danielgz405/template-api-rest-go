package responses

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

// func wil recive w, status (http.Status), message and error

func NoAuthResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorMessage{
		Message: message,
	})
}

func InternalServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorMessage{
		Message: message,
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorMessage{
		Message: message,
	})
}

func NotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorMessage{
		Message: message,
	})
}

func DeleteResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ErrorMessage{
		Message: message,
	})
}
