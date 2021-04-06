package types

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Code    int32       `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(success bool, message string, code int32, data interface{}) *Response {
	return &Response{
		Success: success,
		Message: message,
		Code:    code,
		Data:    data,
	}
}

func NewSuccess(message string, code int32, data interface{}) *Response {
	return NewResponse(true, message, code, data)
}

func NewError(message string, code int32) *Response {
	return NewResponse(false, message, code, nil)
}

func (r *Response) Write(w http.ResponseWriter, code int) {
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
