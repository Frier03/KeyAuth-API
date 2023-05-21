package routes

import (
	"github.com/gin-gonic/gin"

	apiKey "github.com/Frier03/KeyAuth-API/pkg/controllers/api-key"
	apiKeyRole "github.com/Frier03/KeyAuth-API/pkg/controllers/api-key/role"

	"github.com/Frier03/KeyAuth-API/pkg/dependencies"

	middleware "github.com/Frier03/KeyAuth-API/pkg/middleware"
	middlewareApiKey "github.com/Frier03/KeyAuth-API/pkg/middleware/api-key"

	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// Set up auth keys route
func SetupAPIKeyRoutes(r *gin.Engine, deps dependencies.Dependencies) {
	apiKeyRoutes := r.Group("/api-key")
	{
		apiKeyRoutes.POST("/generate",
			middleware.ValidateModel(&models.APIKeyGenerateRequest{}),
			func(c *gin.Context) {
				apiKey.Generate(c, deps)
			},
		)

		apiKeyRoutes.PATCH("/role/adjust",
			middlewareApiKey.ValidateKey(deps.BadgerService),
			middlewareApiKey.TrackKeyUsage(deps.BadgerService),
			func(c *gin.Context) {
				apiKeyRole.Adjust(c, deps.BadgerService)
			},
		)

		apiKeyRoutes.GET("/usage",
			middlewareApiKey.ValidateKey(deps.BadgerService),
			middlewareApiKey.TrackKeyUsage(deps.BadgerService),
			func(c *gin.Context) {
				apiKey.Usage(c)
			},
		)
	}
}
