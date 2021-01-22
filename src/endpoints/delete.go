package endpoints

import (
	"github.com/go-chi/chi"
	"logbyte/src/db"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		write(w, 400, newError("id parameter left empty", nil))
		return
	}

	data, err := db.Delete(id)
	if err != nil {
		write(w, 400, newError("could not find or delete that notification: "+err.Error(), nil))
		return
	}

	write(w, 200, newSuccess("", data))
}
