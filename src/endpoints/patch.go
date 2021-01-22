package endpoints

import (
	"github.com/go-chi/chi"
	"io/ioutil"
	"logbyte/src/db"
	"net/http"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		write(w, 400, newError("id parameter left empty", nil))
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		write(w, 500, newError("an error occurred whilst reading the request body: "+err.Error(), nil))
		return
	}

	data, err := db.Patch(id, bytes)
	if err != nil {
		write(w, 500, newError("error whilst updating database: "+err.Error(), nil))
		return
	}

	write(w, 200, newSuccess("", data))
}
