package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	"github.com/Frier03/KeyAuth-API/pkg/routes"
)

func main() {
	// Create a new Gin Router
	router := gin.Default()

	router.LoadHTMLGlob("../pkg/adminpanel/templates/*.html")

	router.ForwardedByClientIP = true

	router.SetTrustedProxies([]string{
		"127.0.0.1",
		"::1",
		"localhost",
	})

	// Running GIN in "debug" mode. Switch to "release" mode in production.
	gin.SetMode(gin.DebugMode) //gin.ReleaseMode

	// Create dependencies
	deps, err := dependencies.NewDependencies()
	if err != nil {
		return
	}

	// Set up admin panel route
	routes.SetupAdminPanelRoutes(router)

	// Set up api key routes
	routes.SetupAPIKeyRoutes(router, *deps)

	router.Run(":8080")
}
