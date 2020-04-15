package main

import (
	"log"
	"net/http"

	"github.com/mbraeutig/godo-now/api"
)

func main() {
	http.HandleFunc("/index", api.Index)
	http.HandleFunc("/health", api.Health)
	http.HandleFunc("/create", api.Create)
	http.HandleFunc("/read", api.Read)
	http.HandleFunc("/readall", api.ReadAll)
	http.HandleFunc("/delete", api.Delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
