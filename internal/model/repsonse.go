package model

// ApiResponse Standard API response
type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
