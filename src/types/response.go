package types

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(success bool, message string, data interface{}) *Response {
	if success {
		return &Response{
			Success: success,
			Message: message,
			Data:    data,
		}
	}
	return &Response{
		Success: false,
		Error:   message,
		Data:    data,
	}
}
