package api

import (
	"log"
	"net/http"

	"github.com/mbraeutig/godo-now/fauna"
	"github.com/mbraeutig/godo-now/todo"
	"github.com/mbraeutig/godo-now/util"
)

func ReadAll(w http.ResponseWriter, r *http.Request) {
	c := fauna.Client()
	all, err := todo.ReadAll(c)

	if err != nil {
		log.Println(err)
		util.SendNotFound(w, err)
		return
	}

	w.Write([]byte("["))
	for _, todo := range all {
		w.Write([]byte(todo.Text))
		w.Write([]byte("\n"))
	}
	w.Write([]byte("]"))

}
