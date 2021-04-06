package endpoints

import (
	"logbyte/src/db"
	"logbyte/src/types"
	"net/http"
	"strconv"
)

func Batch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 100
	}
	if limit < 1 {
		types.NewError("limit must be greater than zero", 0).Write(w, 400)
		return
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}
	if offset < 0 {
		types.NewError("limit must be greater than or equal to zero", 0).Write(w, 400)
		return
	}

	data, err := db.Paginate(limit, offset)
	if err != nil {
		types.NewError("could not load log entries: "+err.Error(), 0).Write(w, 400)
		return
	}

	types.NewSuccess("", 0, data).Write(w, 200)
}
