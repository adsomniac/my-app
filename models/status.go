package models

// StatusResponse represents the system health check payload
type StatusResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Database string `json:"database"`
}
