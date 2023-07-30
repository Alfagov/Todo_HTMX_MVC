package app

import (
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/db"
	"github.com/gin-gonic/gin"
)

func (a *App) SetupValidationRoutes(engine *gin.Engine) {

	d := a.Database

	validation := engine.Group("/validate")
	{
		validation.POST("/username", validateUsernameHandler(d))
		//validation.POST("/email", validateEmailHandler)
	}
}

func validateUsernameHandler(database db.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")

		exists, err := database.ExistsUserByUsername(username)
		if err != nil {
			c.HTML(200, "username-form-input", gin.H{
				"Error":    err.Error(),
				"Username": username,
			})
			return
		}

		if exists {
			c.HTML(200, "username-form-input", gin.H{
				"Error":    "Username already taken",
				"Username": username,
			})
			return
		}

		c.HTML(200, "username-form-input", gin.H{
			"Username": username,
		})

	}

}
