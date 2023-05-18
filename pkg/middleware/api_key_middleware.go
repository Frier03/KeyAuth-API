package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyValidationMiddleware is a middleware to validate the API key
func APIKeyValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-Api-Key")
		// Check if the API key is valid
		if !isValidAPIKey(apiKey) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
			})
			return
		}

		c.Next()
	}
}

// isValidAPIKey is a function to validate the API key
func isValidAPIKey(apiKey string) bool {
	return apiKey == "a1f44782-6499-4d1e-948d-1ab37ad23b82-9a79bc5f585b0d7002a3400943585372"
}
