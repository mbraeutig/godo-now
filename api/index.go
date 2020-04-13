package api

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
	w.WriteHeader(http.StatusOK)
}
