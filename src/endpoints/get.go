package endpoints

import (
	"github.com/go-chi/chi"
	"logbyte/src/db"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		write(w, 400, newError("id parameter left empty", nil))
		return
	}

	data, err := db.Get(id)
	if err != nil {
		write(w, 400, newError("could not find notification by that id, or the database malfunctioned: "+err.Error(), nil))
		return
	}

	write(w, 200, newSuccess("", data))
}
