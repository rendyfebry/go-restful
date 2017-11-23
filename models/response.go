package models

type ResponseObj struct {
	Error   interface{} `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
