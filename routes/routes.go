package routes

import (
	"github.com/Alfagov/Todo_HTMX_MVC/handlers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", handlers.LoginGetHandler)
	g.POST("/login", handlers.LoginPostHandler)
	g.GET("/", handlers.RootHandler)
}

func PrivateRoutes(g *gin.RouterGroup, h *handlers.ProductsHandler) {

	g.GET("/products", h.RootProducts)
	g.POST("/add/", h.AddProduct)
	g.DELETE("/delete/:id", h.DeleteProduct)
	g.POST("/update/:id", h.UpdateProduct)
	g.GET("/signout", handlers.LogoutGetHandler)

}
