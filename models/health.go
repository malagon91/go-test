package models

// HealthResponse :
type HealthResponse struct {
	ComponentName string `json:"componentName" example:"api-bootstrap-echo"`
	Status        string `json:"status" example:"pass"`
	Version       string `json:"version" example:"1.0.13"`
}
