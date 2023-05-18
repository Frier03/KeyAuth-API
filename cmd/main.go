package main

import (
	"github.com/Frier03/KeyAuth-API/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set ForwardedByClientIP to true to trust the X-Forwarded-For header
	// This ensures that the correct client IP is used in the request context
	r.ForwardedByClientIP = true

	// Set trusted proxies to handle X-Forwarded-For headers
	r.SetTrustedProxies([]string{
		"127.0.0.1", // IPv4 address for localhost
		"::1",       // IPv6 address for localhost
		"localhost", // Hostname for localhost
	})

	// Running GIN in "debug" mode. Switch to "release" mode in production.
	gin.SetMode(gin.DebugMode) //gin.ReleaseMode

	// Set up authentication routes
	routes.SetupAuthRoutes(r)

	r.Run(":8080") // Start the server
}
