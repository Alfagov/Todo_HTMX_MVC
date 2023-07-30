package app

import (
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/jwtHelper"
	"github.com/gin-gonic/gin"
)

func isLoggedIn(jHelp *jwtHelper.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authentication")
		if err != nil || token == "" {
			c.AbortWithStatus(401)
		}

		username, userId, err := jHelp.Validate(token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("username", username)
		c.Set("userid", userId)
		c.Next()
	}
}
