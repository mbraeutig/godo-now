package main

import (
	"log"
	"net/http"

	"github.com/mbraeutig/godo-now/api"
	// "github.com/mbraeutig/godo-now/api"
)

func main() {

	// c := fauna.Client()
	// all, err := todo.ReadAll(c)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("all:%+v", len(all))

	// for _, t := range all {
	// 	log.Println("! ", t)
	// }

	// log.Printf("all:%+v", all[0])

	http.HandleFunc("/index", api.Index)
	http.HandleFunc("/health", api.Health)
	http.HandleFunc("/create", api.Create)
	http.HandleFunc("/delete", api.Delete)
	http.HandleFunc("/read", api.Read)
	http.HandleFunc("/readall", api.ReadAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
