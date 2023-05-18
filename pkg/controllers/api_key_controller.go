package controllers

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Handler for generating an authentication key
func GenerateAPIKeyHandler(c *gin.Context) {
	// Generate an authentication key
	API_KEY := utils.GenerateAPIKey()

	// Return the authentication key in the response
	c.JSON(http.StatusOK, gin.H{
		"API_KEY": API_KEY,
	})
}
