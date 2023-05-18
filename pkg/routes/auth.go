package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/controllers"
	"github.com/Frier03/KeyAuth-API/pkg/middleware"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// SetupAuthRoutes sets up the authentication routes.
func SetupAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login",
			middleware.ValidateModelMiddleware(&models.LoginRequest{}), // Validate login request payload
			controllers.LoginHandler,                                   // Expected resource logic
		)

		authRoutes.POST("/register",
			middleware.ValidateModelMiddleware(&models.RegisterRequest{}), // Validate register request payload
			controllers.LogoutHandler,                                     // Expected resource logic
		)

		authRoutes.POST("/logout",
			controllers.LogoutHandler, // Expected resource logic
		)
	}
}
