package adminpanel

import (
	"fmt"
	"net/http"
	"time"

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

	var apiKeys []models.APIKey
	for _, data := range apiKeyData {
		apiKey := models.APIKey{
			ID:              data["id"].(string),
			SubjectID:       data["subject_id"].(string),
			PermissionLevel: data["permission_level"].(int),
			Usage:           data["usage"].(int),
			Limit:           data["limit"].(int),
			CreatedAt:       data["created_at"].(string),
			ExpiresAt:       data["expires_at"].(string),
			LastUsed:        data["last_used"].(string),
			Active:          data["active"].(bool),
		}
		apiKey.CreatedAt = formatCreatedAt(apiKey.CreatedAt)
		apiKey.ExpiresAt = formatExpiresAt(apiKey.ExpiresAt)
		apiKey.LastUsed = formatUsedAt(apiKey.LastUsed)
		apiKeys = append(apiKeys, apiKey)
	}

	// Render the dashboard.html template with the updated API key data
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"APIKeys": apiKeys,
	})
}

func RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func formatExpiresAt(expiresAt string) string {
	expirationTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", expiresAt)
	if err != nil {
		return "Invalid date"
	}

	remainingDuration := expirationTime.Sub(time.Now())
	if remainingDuration < 0 {
		return "Expired"
	}

	days := int(remainingDuration.Hours() / 24)
	months := days / 30
	days %= 30

	if months > 0 {
		return fmt.Sprintf("%d months left", months)
	}

	return fmt.Sprintf("%d days left", days)
}

func formatUsedAt(usedAt string) string {
	lastUsedTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", usedAt)
	if err != nil {
		return "Invalid date"
	}

	elapsedDuration := time.Since(lastUsedTime)

	if elapsedDuration < time.Minute {
		return "Recently used"
	}

	if elapsedDuration < time.Hour {
		minutes := int(elapsedDuration.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	if elapsedDuration < time.Hour*24 {
		hours := int(elapsedDuration.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	}

	days := int(elapsedDuration.Hours() / 24)
	if days < 30 {
		return fmt.Sprintf("%d days ago", days)
	}

	months := days / 30
	return fmt.Sprintf("%d months ago", months)
}
func formatCreatedAt(createdAt string) string {
	timestamp, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", createdAt)
	if err != nil {
		return "Invalid date"
	}

	formattedDate := timestamp.Format("2006/01/02")
	return formattedDate
}
