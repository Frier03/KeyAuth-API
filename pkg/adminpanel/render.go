package adminpanel

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
