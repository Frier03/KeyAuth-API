package main

import (
	"github.com/Frier03/KeyAuth-API/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up authentication routes
	routes.SetupAuthRoutes(r)

	r.Run() // Start the server
}
