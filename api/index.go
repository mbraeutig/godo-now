package api

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
	w.WriteHeader(http.StatusOK)
}

func TestFauna(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("FAUNADB_SECRET")
	log.Println("Secret: " + secret)
}
