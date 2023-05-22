package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Frier03/KeyAuth-API/pkg/adminpanel"
	controllers "github.com/Frier03/KeyAuth-API/pkg/controllers/adminpanel"
	middleware "github.com/Frier03/KeyAuth-API/pkg/middleware/adminpanel"
)

// Set up auth keys route
func SetupAdminPanelRoutes(r *gin.Engine) {
	adminPanelRoutes := r.Group("/adminpanel")
	{
		adminPanelRoutes.GET("/",
			middleware.Authorization(),
			adminpanel.RenderIndexPage,
		)

		adminPanelRoutes.GET("/login", adminpanel.RenderLoginPage)
		adminPanelRoutes.POST("/login", controllers.ProcessLogin)
	}
}
