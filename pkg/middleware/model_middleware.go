package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/models"
)

func ValidateModelMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request payload",
			})
			return
		}

		// Type assertion to access the specific model fields
		switch m := model.(type) {
		case *models.LoginRequest:
			// Check if username and password are provided
			if m.Username == "" || m.Password == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "Missing required fields",
				})
				return
			}

		case *models.RegisterRequest:
			// Check if username and password are provided
			if m.Username == "" || m.Password == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "Missing required fields",
				})
				return
			}

		// Add cases for other models as needed

		default:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Set("model", model)
		c.Next()
	}
}
