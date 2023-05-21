package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/controllers"
	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	"github.com/Frier03/KeyAuth-API/pkg/middleware"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// Set up auth keys route
func SetupAPIKeyRoutes(r *gin.Engine, deps dependencies.Dependencies) {
	apiKeyRoutes := r.Group("/api-key")
	{
		apiKeyRoutes.GET("/generate",
			middleware.ValidateModelMiddleware(&models.APIKeyGenerateRequest{}),
			func(c *gin.Context) {
				controllers.GenerateAPIKeyHandler(c, deps)
			},
		)

		apiKeyRoutes.GET("/usage",
			middleware.APIKeyValidationMiddleware(),
			func(c *gin.Context) {
				controllers.APIKeyUsageHandler(c, deps)
			},
		)
	}
}
