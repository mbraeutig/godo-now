package api

import (
	"log"
	"net/http"

	"github.com/mbraeutig/godo-now/fauna"
	"github.com/mbraeutig/godo-now/todo"
	"github.com/mbraeutig/godo-now/util"
)

func ReadAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	all, err := todo.ReadAll(ctx, firestore.Client)

	if err != nil {
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
