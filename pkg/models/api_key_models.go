package models

// APIKeyGenerateRequest represents the request body for the register handler
type APIKeyGenerateRequest struct {
	UUID string `json:"uuid"`
}

// APIKey struct that represents the API key table in the database
type APIKey struct {
	ID              string `json:"id"`
	KEY             string `json:"key"`
	SubjectID       string `json:"subject_id"`
	PermissionLevel int    `json:"permission_level"`
	Usage           int    `json:"usage"`
	Limit           int    `json:"limit"`
	CreatedAt       string `json:"created_at"`
	ExpiresAt       string `json:"expires_at"`
	LastUsed        string `json:"last_used"`
	Active          bool   `json:"active"`
}
