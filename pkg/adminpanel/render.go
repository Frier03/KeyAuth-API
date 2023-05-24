package adminpanel

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/services"
	"github.com/gin-gonic/gin"
)

func RenderDashboardPage(c *gin.Context, badgerService *services.BadgerService) {

	apiKeys, err := badgerService.GetAllData()
	if err != nil {
		// Handle error
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	totalGeneratedAPIKeys := badgerService.FetchTotalGeneratedAPIKeys()
	totalExpiredAPIKeys := badgerService.FetchTotalExpiredAPIKeys()

	// Render the dashboard.html template with the updated API key data
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"APIKeys":               apiKeys,
		"TotalGeneratedAPIKeys": totalGeneratedAPIKeys,
		"TotalExpiredAPIKeys":   totalExpiredAPIKeys,
	})
}

func RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
