package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *App) SetupPagesRoutes(engine *gin.Engine) {

	pages := engine.Group("/", isLoggedIn(a.JWT))
	{
		pages.GET("/", enforceLoggedIn(), indexPageHandler)
		pages.GET("/login", enforceLoggedOut(), loginPageHandler)
		pages.GET("/register", enforceLoggedOut(), registerPageHandler)
	}
}

func indexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"IsAuthenticated": c.GetBool("logged_in")})
}

func loginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func registerPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
