package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Default role access",
	})
}
