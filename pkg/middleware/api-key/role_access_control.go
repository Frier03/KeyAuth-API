package middleware

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/gin-gonic/gin"
)

func RoleAccessControl(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyData, _ := c.Get("api-key-model")

		// Apply model to apiKeyData
		apiKeyModel, _ := apiKeyData.(models.APIKey)

		// Get api key permission level
		apiPermissionLevel := models.Role(apiKeyModel.PermissionLevel)

		// Get requiredRole permission level
		requiredPermissionLevel := role

		// Check if APIPermissionLevel has less privileges than the required role
		if apiPermissionLevel < requiredPermissionLevel {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Insufficient privileges to access this resource.",
			})
			return
		}

		c.Next()
	}
}
