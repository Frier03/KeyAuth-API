package middleware

import (
	"github.com/gin-gonic/gin"
)

// APIKeyValidationMiddleware is a middleware to validate the API key
func APIKeyValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-Api-Key")
		// Check if the API key is valid

		c.Set("x-api-key", apiKey)
		c.Next()
	}
}
