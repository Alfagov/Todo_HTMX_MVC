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
		users.POST("/register", validateUserWithDB(d), registerUserHandler(d))
		users.POST("/login", loginUserHandler(d, a.JWT))
		users.GET("/logout", logoutHandler)
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

func validateUserWithDB(database db.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		name := c.PostForm("name")
		password := c.PostForm("password")

		exists, err := database.ExistsUserByUsername(username)
		if err != nil {
			c.Header("HX-Retarget", "#submit-w-error")
			c.HTML(http.StatusOK, "submit-with-error", gin.H{
				"ServerError": err.Error(),
			})
			return
		}

		if exists {
			c.Header("HX-Retarget", "#username-field")
			c.HTML(200, "username-form-input", gin.H{
				"Error":    "Username already taken",
				"Username": username,
			})
			return
		}

		exists, err = database.ExistsUserByEmail(email)
		if err != nil {
			c.Header("HX-Retarget", "#submit-w-error")
			c.HTML(http.StatusOK, "submit-with-error", gin.H{
				"ServerError": err.Error(),
			})
			return
		}

		if exists {
			c.Header("HX-Retarget", "#email-field")
			c.HTML(200, "email-form-input", gin.H{
				"Error": "Email already taken",
				"Email": email,
			})
			return
		}

		c.Set("validuser", models.CreateUser(username, name, email, password))
		c.Next()
	}
}

func registerUserHandler(database db.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, ok := c.Get("validuser")
		if !ok {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"title": "Register Page",
			})
			return
		}

		err := database.CreateUser(user.(*models.User))
		if err != nil {
			c.Header("HX-Retarget", "#submit-w-error")
			c.HTML(http.StatusOK, "submit-with-error", gin.H{
				"ServerError": err.Error(),
			})
			return
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
