package notifs

import (
	"encoding/json"
	"log"
	"net/http"
	"notifs/src/types"
)

func write(w http.ResponseWriter, status int, data *types.Response) {
	str, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(status)
	w.Write(str)
}

func newSuccess(message string, data interface{}) *types.Response {
	return types.NewResponse(true, message, data)
}

func newError(message string, data interface{}) *types.Response {
	return types.NewResponse(false, message, data)
}
