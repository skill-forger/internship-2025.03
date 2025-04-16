package contract

// HealthCheckResponse specifies the data and types for health check API response
type HealthCheckResponse struct {
	Resource string `json:"resource,omitempty"`
	Status   string `json:"status,omitempty"`
}
