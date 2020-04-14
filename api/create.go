package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mbraeutig/godo-now/fauna"
	"github.com/mbraeutig/godo-now/todo"
	"github.com/mbraeutig/godo-now/util"
)

func decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var t todo.Todo
	if err := decode(r, &t); err != nil {
		log.Println(err)
		util.SendInternalServerError(w, err)
		return
	}

	c := fauna.Client()

	id, err := todo.Create(c, t)
	if err != nil {
		log.Println(err)
		util.SendInternalServerError(w, err)
		return
	}

	type viewTodo struct {
		Text     string
		Complete bool
		ID       string
	}

	util.SendJSON(w, viewTodo{
		Text:     t.Text,
		Complete: t.Complete,
		ID:       id,
	})
}
