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
		types.NewError("malformed json body", 0).Write(w, 400)
		return
	}

	data.Id = random.MustSecureString(12, hex)
	data.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.Insert(data)
	if err != nil {
		types.NewError("could not insert document: "+err.Error(), 0).Write(w, 500)
		return
	}

	types.NewSuccess("created notification", 0, data).Write(w, 200)
}
