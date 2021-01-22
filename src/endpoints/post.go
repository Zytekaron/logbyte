package endpoints

import (
	"encoding/json"
	"github.com/zytekaron/gotil/random"
	"logbyte/src/db"
	"logbyte/src/types"
	"net/http"
	"time"
)

var hex = []rune("0123456789abcdef")

func Post(w http.ResponseWriter, r *http.Request) {
	var data *types.LogEntry
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		write(w, 400, newError("malformed json body", nil))
		return
	}

	data.Id = random.MustSecureString(12, hex)
	data.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.Insert(data)
	if err != nil {
		write(w, 500, newError("could not insert document: "+err.Error(), nil))
		return
	}

	type Id struct {
		Id string `json:"id"`
	}
	write(w, 200, newSuccess("created notification", &Id{Id: data.Id}))
}
