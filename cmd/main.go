package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up authentication routes
	routes.setupAuthRoutes(r)

	r.Run() // Start the server
}
