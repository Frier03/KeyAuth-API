package routes

import (
	"github.com/gin-gonic/gin"

	auth "github.com/Frier03/KeyAuth-API/pkg/controllers/auth"
	"github.com/Frier03/KeyAuth-API/pkg/middleware"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// SetupAuthRoutes sets up the authentication routes.
func SetupAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login",
			middleware.ValidateModelMiddleware(&models.LoginRequest{}), // Validate login request payload
			auth.Login, // Expected resource logic
		)

		authRoutes.POST("/register",
			middleware.ValidateModelMiddleware(&models.RegisterRequest{}), // Validate register request payload
			auth.Register, // Expected resource logic
		)

		authRoutes.POST("/logout",
			auth.Logout, // Expected resource logic
		)
	}
}
