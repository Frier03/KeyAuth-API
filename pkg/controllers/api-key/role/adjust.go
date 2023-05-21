package role

import (
	"net/http"
	"strconv"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/Frier03/KeyAuth-API/pkg/services"
	"github.com/gin-gonic/gin"
)

func Adjust(c *gin.Context, badgerService *services.BadgerService) {
	apiKey := c.GetHeader("X-Api-Key")
	adjust, _ := c.GetQuery("permission_level")
	apiKeyData, _ := c.Get("api-key-model")

	// Apply model to apiKey
	apiKeyModel, _ := apiKeyData.(models.APIKey)

	// Validate adjust value
	level, err := strconv.Atoi(adjust)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid adjust value",
		})
		return
	}

	// Check if level is within the allowed range
	if level < int(models.Default) || level > int(models.Admin) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Adjust value is out of range",
		})
		return
	}

	apiKeyModel.PermissionLevel = level

	// Update the API key in the database
	err = badgerService.PutAPIKey([]byte(apiKey), &apiKeyModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update API key",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
