package endpoints

import (
	"logbyte/src/db"
	"logbyte/src/types"
	"net/http"
)

func Count(w http.ResponseWriter, r *http.Request) {
	count, err := db.Count()
	if err != nil {
		types.NewError("database malfunction: "+err.Error(), 0).Write(w, 500)
		return
	}

	types.NewSuccess("", 0, count).Write(w, 200)
}
