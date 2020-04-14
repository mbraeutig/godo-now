package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/mbraeutig/godo-now/fauna"
	"github.com/mbraeutig/godo-now/todo"
	"github.com/mbraeutig/godo-now/util"
)

func Read(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		err := errors.New("value 'id' not in the query")
		log.Println(err)
		util.SendInternalServerError(w, err)
		return
	}

	c := fauna.Client()
	t, err := todo.Get(c, id)

	if err != nil {
		log.Println(err)
		util.SendNotFound(w, err)
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
