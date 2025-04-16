package models

// APIResponse represents a standard API response structure
type APIResponse struct {
	Status  int         `json:"status"`           // HTTP status code
	Data    interface{} `json:"data,omitempty"`   // Response data
	Message string      `json:"message"`          // Message for client
}
