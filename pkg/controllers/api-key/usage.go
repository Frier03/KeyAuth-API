/*
/api-key/usage
*/

package controllers

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/gin-gonic/gin"
)

func Usage(c *gin.Context, deps dependencies.Dependencies) {
	apiKeyData, _ := c.Get("api-key-model")

	// Apply model to apiKey
	apiKeyModel, _ := apiKeyData.(models.APIKey)

	c.JSON(http.StatusOK, gin.H{
		"message": apiKeyModel,
	})
}
