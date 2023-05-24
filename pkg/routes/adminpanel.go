package routes

import (
	"github.com/Frier03/KeyAuth-API/pkg/adminpanel"
	controllers "github.com/Frier03/KeyAuth-API/pkg/controllers/adminpanel"
	"github.com/Frier03/KeyAuth-API/pkg/dependencies"
	middleware "github.com/Frier03/KeyAuth-API/pkg/middleware/adminpanel"
	"github.com/gin-gonic/gin"
)

// Set up auth keys route
func SetupAdminPanelRoutes(r *gin.Engine, deps dependencies.Dependencies) {
	adminPanelRoutes := r.Group("/adminpanel")
	{
		adminPanelRoutes.GET("/",
			middleware.Authorization(),
			func(c *gin.Context) {
				adminpanel.RenderDashboardPage(c, deps.BadgerService)
			},
		)

		adminPanelRoutes.GET("/login", adminpanel.RenderLoginPage)
		adminPanelRoutes.POST("/login", controllers.ProcessLogin)
	}
}
