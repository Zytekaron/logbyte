package endpoints

import (
	"logbyte/src/db"
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
		write(w, 400, newError("limit must be greater than zero", nil))
		return
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}
	if offset < 0 {
		write(w, 400, newError("limit must be greater than or equal to zero", nil))
		return
	}

	data, err := db.Paginate(limit, offset)
	if err != nil {
		write(w, 400, newError("could not load log entries: "+err.Error(), nil))
		return
	}

	write(w, 200, newSuccess("", data))
}
