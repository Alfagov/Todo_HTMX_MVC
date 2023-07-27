package handlers

import (
	"log"
	"net/http"

	"github.com/Alfagov/Todo_HTMX_MVC/globals"
	"github.com/Alfagov/Todo_HTMX_MVC/helpers"
	"github.com/Alfagov/Todo_HTMX_MVC/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {

	// get the navbar links
	nLinks := models.NewNavBarLinks(models.DefaultNavBarLinks)

	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":   "Home Page",
		"Columns": 3,
		"Links":   nLinks.Links,
	})
}

func LoginGetHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	// if user is already logged in, redirect to main page
	if user != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"user": user})
		return
	}

	// if user is not logged in, show login page
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
		"user":    user,
	})
}

func LoginPostHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user != nil {
		c.HTML(http.StatusBadRequest, "main.html", gin.H{"user": user})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if helpers.EmptyUserPass(username, password) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
		return
	}

	if !helpers.CheckUserPass(username, password) {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
		return
	}

	session.Set(globals.Userkey, username)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/products")
}

func LogoutGetHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	log.Println("logging out user:", user)
	if user == nil {
		log.Println("Invalid session token")
		return
	}
	session.Delete(globals.Userkey)
	if err := session.Save(); err != nil {
		log.Println("Failed to save session:", err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
