package controllers

import (
	"net/http"

	"github.com/Frier03/KeyAuth-API/pkg/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	// Retrieve the login request model from the context
	req, _ := ctx.Get("model")

	// Try assert the request model to *models.LoginRequest
	loginReq, _ := req.(*models.LoginRequest)

	// Access the username and password from the request
	_, _ = loginReq.Username, loginReq.Username // Intentionally ignoring these variables

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func RegisterHandler(ctx *gin.Context) {
	// Retrieve the login request model from the context
	req, _ := ctx.Get("model")

	// Try assert the request model to *models.LoginRequest
	registerReq, _ := req.(*models.RegisterRequest)

	// Access the username and password from the request
	_, _ = registerReq.Username, registerReq.Username // Intentionally ignoring these variables

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
}

func LogoutHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
