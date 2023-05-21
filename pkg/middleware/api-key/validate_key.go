package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/Frier03/KeyAuth-API/pkg/services"
	"github.com/gin-gonic/gin"
)

func ValidateKey(badgerService *services.BadgerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Api-Key")

		// Lookup the API key in the database using badgerService
		value, err := badgerService.Get([]byte(apiKey))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": "Invalid API key",
			})
			return
		}

		// Is the API Key not in the database?
		if value == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "API key Not Found",
			})
			return
		}

		// Unmarshal the API key value into an APIKey object
		var apiKeyModel models.APIKey
		err = json.Unmarshal(value, &apiKeyModel)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "API key has invalid format?",
			})
			return
		}

		// Is the API Key not active?
		if !apiKeyModel.Active {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Inactive API key",
			})
			return
		}

		c.Set("api-key-model", apiKeyModel)
		c.Next()
	}
}
