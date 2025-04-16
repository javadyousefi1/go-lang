package models

// @Description This is all api response structure
type APIResponse struct {
	Status  int         `json:"status"`           // HTTP status code
	Data    interface{} `json:"data,omitempty"`   // Response data
	Message string      `json:"message"`          // Message for client
}
