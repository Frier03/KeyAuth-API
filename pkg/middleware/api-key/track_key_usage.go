package middleware

import (
	"github.com/Frier03/KeyAuth-API/pkg/services"
	"github.com/gin-gonic/gin"
)

func TrackKeyUsage(badgerService *services.BadgerService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
