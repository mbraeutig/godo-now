package api

import (
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("FAUNADB_SECRET")
	io.WriteString(w, "Secret: "+secret)
}
