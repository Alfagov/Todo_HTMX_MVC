package app

import (
	"github.com/Alfagov/Todo_HTMX_MVC/models"
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/db"
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/jwtHelper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (a *App) SetupUserRoutes(engine *gin.Engine) {

	d := a.Database

	users := engine.Group("/users")
	{
		users.POST("/register", registerUserHandler(d))
		users.POST("/login", loginUserHandler(d, a.JWT))
		users.POST("/logout", logoutHandler)
		/*users.GET("/me", me)
		users.GET("/userByName/:name", userByName)
		users.GET("/userByUsername/:username", userByUsername)*/
	}

}

func logoutHandler(c *gin.Context) {
	c.SetCookie("Authentication", "", -1, "/", "localhost", false, true)
	c.Header("HX-Redirect", "/login/")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func registerUserHandler(database db.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		name := c.PostForm("name")

		exists, err := database.ExistsUserByUsername(username)
		if err != nil {
			log.Println(err)
		}

		if exists {
			log.Println("User already exists")
		}

		user := models.CreateUser(username, name, email, password)

		err = database.CreateUser(user)
		if err != nil {
			log.Println(err)
		}

		c.Header("HX-Redirect", "/login/")
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login Page",
		})
	}

}

func loginUserHandler(database db.Dao, jHelp *jwtHelper.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user, err := database.LoginUser(username, password)
		if err != nil {
			log.Println("LOGIN ERROR", err)
		}

		token, err := jHelp.Create(time.Duration(24)*time.Hour, user.Username, user.Id)
		if err != nil {
			log.Println("LOGIN ERROR", err)
		}

		c.SetCookie("Authentication", token, 3600, "/", "localhost", false, true)
		c.Header("HX-Redirect", "/")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IsAuthenticated": true,
		})
	}
}
