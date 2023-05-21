package middleware

import (
	"net/http"
	"time"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/Frier03/KeyAuth-API/pkg/services"
	"github.com/gin-gonic/gin"
)

func TrackKeyUsage(badgerService *services.BadgerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Api-Key")
		apiKeyData, _ := c.Get("api-key-model")

		// Apply model to apiKeyData
		apiKeyModel, _ := apiKeyData.(models.APIKey)

		// Check if limit on API key has exceeded
		if apiKeyModel.Limit <= apiKeyModel.Usage {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}

		// Increment the API key usage count
		apiKeyModel.Usage++

		// Update the API last used
		apiKeyModel.LastUsed = time.Now()

		// Update the API key in the database
		err := badgerService.PutAPIKey([]byte(apiKey), &apiKeyModel)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update API key",
			})
			return
		}

		// Update c set for api-key-model
		c.Set("api-key-model", apiKeyModel)
		c.Next()
	}
}
