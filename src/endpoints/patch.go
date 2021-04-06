package endpoints

import (
	"github.com/go-chi/chi"
	"io/ioutil"
	"logbyte/src/db"
	"logbyte/src/types"
	"net/http"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		types.NewError("id parameter left empty", 0).Write(w, 400)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		types.NewError("an error occurred whilst reading the request body: "+err.Error(), 0).Write(w, 500)
		return
	}

	data, err := db.Patch(id, bytes)
	if err != nil {
		types.NewError("error whilst updating database: "+err.Error(), 0).Write(w, 500)
		return
	}

	types.NewSuccess("", 0, data).Write(w, 200)
}
