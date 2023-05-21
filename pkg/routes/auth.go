package routes

import (
	"github.com/gin-gonic/gin"

	authController "github.com/Frier03/KeyAuth-API/pkg/controllers/auth"
	middleware "github.com/Frier03/KeyAuth-API/pkg/middleware"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// SetupAuthRoutes sets up the authentication routes.
func SetupAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login",
			middleware.ValidateModel(&models.LoginRequest{}), // Validate login request payload
			authController.Login,                             // Expected resource logic
		)

		authRoutes.POST("/register",
			middleware.ValidateModel(&models.RegisterRequest{}), // Validate register request payload
			authController.Register,                             // Expected resource logic
		)

		authRoutes.POST("/logout",
			authController.Logout, // Expected resource logic
		)
	}
}
