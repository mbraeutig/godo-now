package util

import (
	"encoding/json"
	"net/http"
)

// SendJSON - send json with 200 status
func SendJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
}

// SendNotFound - send 404 error with error in message key
func SendNotFound(w http.ResponseWriter, err error) {
	sendError(w, err, http.StatusNotFound)
}

// SendSuccess - send 200 message with empyt json body
func SendSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
	w.WriteHeader(http.StatusOK)
}

// SendInternalServerError - send 500 error with error in message key
func SendInternalServerError(w http.ResponseWriter, err error) {
	sendError(w, err, http.StatusInternalServerError)
}

func sendError(w http.ResponseWriter, err error, code int) {
	data := map[string]string{
		"message": err.Error(),
	}

	if js, _ := json.Marshal(data); js != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
