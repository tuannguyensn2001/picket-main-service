package app

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}
