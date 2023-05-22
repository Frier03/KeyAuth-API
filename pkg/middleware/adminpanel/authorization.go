package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/Frier03/KeyAuth-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")

		if err != nil {
			c.Redirect(http.StatusSeeOther, "/adminpanel/login")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/adminpanel/login")
			c.Abort()
			return
		}

		// Add the claims to request
		c.Set("claims", claims)

		// Get username from claims
		username := claims["sub"].(string)

		// Generate new jwt and set in cookie
		token, err = utils.GenerateToken(username, 30*time.Minute)
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

		c.Next()
	}
}
