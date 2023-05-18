package models

// LoginRequest represents the request body for the login handler
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents the request body for the register handler
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
