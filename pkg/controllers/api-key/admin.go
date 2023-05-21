package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin role access",
	})
}
