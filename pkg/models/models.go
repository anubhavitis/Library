package models

type Response struct {
	Success string `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}
