package api

import (
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_ = os.Getenv("FAUNADB_SECRET")
	// io.WriteString(w, "Secret: "+secret)
}
