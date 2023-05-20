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
	ID        string    `json:"id"`
	APIKey    string    `json:"api_key"`
	Usage     int       `json:"usage"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
