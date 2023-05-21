package routes

import (
	"github.com/gin-gonic/gin"

	apiController "github.com/Frier03/KeyAuth-API/pkg/controllers/api-key"
	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	middleware "github.com/Frier03/KeyAuth-API/pkg/middleware"
	apiMiddleware "github.com/Frier03/KeyAuth-API/pkg/middleware/api-key"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// Set up auth keys route
func SetupAPIKeyRoutes(r *gin.Engine, deps dependencies.Dependencies) {
	apiKeyRoutes := r.Group("/api-key")
	{
		apiKeyRoutes.POST("/generate",
			middleware.ValidateModel(&models.APIKeyGenerateRequest{}),
			func(c *gin.Context) {
				apiController.Generate(c, deps)
			},
		)

		apiKeyRoutes.GET("/usage",
			apiMiddleware.ValidateKey(deps.BadgerService),
			func(c *gin.Context) {
				apiController.Usage(c, deps)
			},
		)
	}
}
