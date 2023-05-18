package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/controllers"
)

// Set up auth keys route
func SetupAPIKeyRoutes(r *gin.Engine) {
	apiKeyRoutes := r.Group("/api-key")
	{
		apiKeyRoutes.GET("/generate", controllers.GenerateAPIKeyHandler)
	}
}
