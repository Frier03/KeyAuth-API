package models

import (
	"time"
)

// APIKeyGenerateRequest represents the request body for the register handler
type APIKeyGenerateRequest struct {
	UUID string `json:"uuid"`
}

// APIKey struct that represents the API key table in the database
type APIKey struct {
	ID              string    `json:"id"`
	SubjectID       string    `json:"subject_id"`
	PermissionLevel int       `json:"permission_level"`
	Usage           int       `json:"usage"`
	Limit           int       `json:"limit"`
	CreatedAt       time.Time `json:"created_at"`
	ExpiresAt       time.Time `json:"expires_at"`
	LastUsed        time.Time `json:"last_used"`
	Active          bool      `json:"active"`
}
