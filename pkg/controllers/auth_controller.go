package controllers

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
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

func RegisterHandler(c *gin.Context) {
	// Retrieve the login request model from the context
	req, _ := c.Get("model")

	// Try assert the request model to *models.LoginRequest
	registerReq, _ := req.(*models.RegisterRequest)

	// Access the username and password from the request
	_, _ = registerReq.Username, registerReq.Username // Intentionally ignoring these variables

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
}

func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
