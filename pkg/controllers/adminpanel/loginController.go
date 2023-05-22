package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Frier03/KeyAuth-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ProcessLogin(c *gin.Context) {
	// Get the login form data
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate login credentials

	// Placeholder for validation
	if username != "admin" || password != "123" {
		// Invalid credentials, show an error message
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": "Invalid credentials",
		})

		return
	}

	// Generate JWT
	token, err := utils.GenerateToken(username, 30*time.Minute)
	if err != nil {
		log.Fatal("Error generating JWT token", err)
	}

	// Set the JWT in cookies
	c.SetCookie(
		"jwt",
		token,
		int(30*time.Minute.Seconds()),
		"/adminpanel",
		"",
		false,
		true,
	)

	// Redirect to the admin panel
	c.Redirect(http.StatusSeeOther, "/adminpanel")
}
