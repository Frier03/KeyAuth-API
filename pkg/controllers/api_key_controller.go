package controllers

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Handler for generating an authentication key
func GenerateAPIKeyHandler(c *gin.Context) {
	// Generate an authentication key
	apiKey := utils.GenerateAPIKey()

	// Return the authentication key in the response
	c.JSON(http.StatusOK, gin.H{
		"API_KEY": apiKey,
	})
}

// Handler for the /api-key/usage endpoint
func APIKeyUsageHandler(c *gin.Context) {
	// Your logic to check the API key usage here

	// Assuming you want to return a JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "API key usage endpoint",
	})
}
