package endpoints

import (
	"github.com/go-chi/chi"
	"logbyte/src/db"
	"logbyte/src/types"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		types.NewError("id parameter left empty", 0).Write(w, 400)
		return
	}

	data, err := db.Delete(id)
	if err != nil {
		types.NewError("could not find or delete that notification: "+err.Error(), 0).Write(w, 400)
		return
	}

	types.NewSuccess("", 0, data).Write(w, 200)
}
