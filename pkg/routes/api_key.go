package routes

import (
	"github.com/gin-gonic/gin"

	api "github.com/Frier03/KeyAuth-API/pkg/controllers/api-key"
	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	"github.com/Frier03/KeyAuth-API/pkg/middleware"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// Set up auth keys route
func SetupAPIKeyRoutes(r *gin.Engine, deps dependencies.Dependencies) {
	apiKeyRoutes := r.Group("/api-key")
	{
		apiKeyRoutes.POST("/generate",
			middleware.ValidateModelMiddleware(&models.APIKeyGenerateRequest{}),
			func(c *gin.Context) {
				api.Generate(c, deps)
			},
		)

		apiKeyRoutes.GET("/usage",
			middleware.APIKeyValidationMiddleware(deps.BadgerService),
			func(c *gin.Context) {
				api.Usage(c, deps)
			},
		)
	}
}
