/*
/api-key/generate
*/

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

func Generate(c *gin.Context, deps dependencies.Dependencies) {
	// Retrieve the generate api key request model from the context
	req, _ := c.Get("model")

	// Apply model to req
	reqModel, _ := req.(*models.APIKeyGenerateRequest)

	// Get UUID from request
	uuid := reqModel.UUID

	// Generate a new UUID for the api key
	ID := utils.GenerateUUID()

	// Generate an authentication key
	apiKey := utils.GenerateAPIKey()

	// Create an instance of the APIKey struct
	apiKeyInstance := &models.APIKey{
		ID:        ID,
		Usage:     0,
		Limit:     1000,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().AddDate(0, 2, 1), // Example expires at time, two months and 1 day from now
		LastUsed:  time.Now(),
		Active:    true,
		SubjectID: uuid,
		Roles:     []string{"default"},
	}

	// Store the APIKey in the database
	err := deps.BadgerService.PutAPIKey([]byte(apiKey), apiKeyInstance)
	if err != nil {
		log.Fatal(err)
	}

	// Return the authentication key in the response
	c.JSON(http.StatusCreated, gin.H{
		"API_KEY": apiKey,
	})
}
