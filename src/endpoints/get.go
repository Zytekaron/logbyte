package endpoints

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"logbyte/src/db"
	"logbyte/src/types"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		types.NewError("id parameter left empty", 0).Write(w, 400)
		return
	}

	data, err := db.Get(id)
	if err == mongo.ErrNoDocuments {
		types.NewError("could not find notification by that id", 0).Write(w, 400)
	}
	if err != nil {
		types.NewError("database malfunction: "+err.Error(), 0).Write(w, 500)
		return
	}

	types.NewSuccess("", 0, data).Write(w, 200)
}
