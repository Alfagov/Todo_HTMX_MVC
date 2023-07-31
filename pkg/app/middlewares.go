package app

import (
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/jwtHelper"
	"github.com/gin-gonic/gin"
)

func isLoggedIn(jHelp *jwtHelper.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authentication")
		if err != nil || token == "" {
			c.Set("logged_in", false)
			c.Next()
			return
		}

		username, userId, err := jHelp.Validate(token)
		if err != nil {
			c.Set("logged_in", false)
			c.Next()
			return
		}

		c.Set("logged_in", true)
		c.Set("username", username)
		c.Set("userid", userId)
		c.Next()
	}
}

func enforceLoggedOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool("logged_in") {
			c.Redirect(302, "/")
			c.Abort()
			return
		}

		c.Next()
	}
}

func enforceLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !c.GetBool("logged_in") {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
