package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/middleware"
	"github.com/Frier03/KeyAuth-API/pkg/models"
)

// SetupAuthRoutes sets up the authentication routes.
func SetupAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login",
			middleware.ValidateModelMiddleware(&models.LoginRequest{}), // Validate login request payload
			loginHandler, // Expected resource logic
		)

		authRoutes.POST("/register",
			middleware.ValidateModelMiddleware(&models.RegisterRequest{}), // Validate register request payload
			registerHandler, // Expected resource logic
		)

		authRoutes.POST("/logout",
			logoutHandler, // Expected resource logic
		)
	}
}

func loginHandler(c *gin.Context) {
	// Retrieve the login request model from the context
	req, _ := c.Get("model")

	// Try assert the request model to *models.LoginRequest
	loginReq, _ := req.(*models.LoginRequest)

	// Access the username and password from the request
	_, _ = loginReq.Username, loginReq.Username // Intentionally ignoring these variables

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func registerHandler(c *gin.Context) {
	// Retrieve the login request model from the context
	req, _ := c.Get("model")

	// Try assert the request model to *models.LoginRequest
	registerReq, _ := req.(*models.RegisterRequest)

	// Access the username and password from the request
	_, _ = registerReq.Username, registerReq.Username // Intentionally ignoring these variables

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func logoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout successful",
	})
}
