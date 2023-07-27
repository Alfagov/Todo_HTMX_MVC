package main

import (
	"log"

	"github.com/Alfagov/Todo_HTMX_MVC/globals"
	"github.com/Alfagov/Todo_HTMX_MVC/handlers"
	"github.com/Alfagov/Todo_HTMX_MVC/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// create new todoHandler
	h := handlers.NewTodoHandler()

	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")
	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	routes.PublicRoutes(public)

	private := r.Group("/")
	routes.PrivateRoutes(private, &h)

	log.Fatal(r.Run(":8080"))

}
