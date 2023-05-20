package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/Frier03/KeyAuth-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Handler for generating an authentication key
func GenerateAPIKeyHandler(c *gin.Context, deps dependencies.Dependencies) {
	// Retrieve the generate api key request model from the context
	req, _ := c.Get("model")

	// Try assert the request model to *models.APIKEYGenerateRequest
	genReq, _ := req.(*models.APIKeyGenerateRequest)

	// Get UUID
	uuid := genReq.UUID

	// Generate an authentication key
	apiKey := utils.GenerateAPIKey()

	// Create an instance of the APIKey struct
	apiKeyInstance := &models.APIKey{
		ID:        uuid,
		APIKey:    apiKey,
		Usage:     0,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour),
	}

	// Store the APIKey in the database
	err := deps.BadgerService.PutAPIKey([]byte("api_key"), apiKeyInstance)
	if err != nil {
		log.Fatal(err)
	}

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
