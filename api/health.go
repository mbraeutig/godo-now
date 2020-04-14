package api

import (
	"net/http"

	"github.com/mbraeutig/godo-now/util"
)

func Health(w http.ResponseWriter, r *http.Request) {
	util.SendSuccess(w)
}
