package adminpanel

import (
	"encoding/json"
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/Frier03/KeyAuth-API/pkg/services"
	"github.com/gin-gonic/gin"
)

func RenderDashboardPage(c *gin.Context, badgerService *services.BadgerService) {
	apiKeyData, err := badgerService.GetAllData()
	if err != nil {
		// Handle error
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Convert the map data to models.APIKey struct
	var apiKeys []models.APIKey

	for _, data := range apiKeyData {
		apiKeyBytes, err := json.Marshal(data)
		if err != nil {
			// Handle error
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var apiKey models.APIKey
		err = json.Unmarshal(apiKeyBytes, &apiKey)
		if err != nil {
			// Handle error
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		apiKeys = append(apiKeys, apiKey)
	}

	// Render the dashboard.html template with the API key data
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"APIKeys": apiKeys,
	})
}

func RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
