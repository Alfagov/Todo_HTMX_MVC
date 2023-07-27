package handlers

import (
	"net/http"

	"github.com/Alfagov/Todo_HTMX_MVC/models"
	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	ProductsList []models.Product
}

func NewTodoHandler() ProductsHandler {
	return ProductsHandler{models.DefaultProductList}
}

func (t *ProductsHandler) RootProducts(c *gin.Context) {

	nLinks := models.NewNavBarLinks(models.DefaultNavBarLinks)

	c.HTML(http.StatusOK, "products.html", gin.H{
		"title":    "Products page",
		"Products": t.ProductsList,
		"Columns":  3,
		"Links":    nLinks.Links,
	})
}

func (t *ProductsHandler) AddProduct(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	newElem := models.NewProduct(name, "Category 1", 100, 10, desc, 1, "https://picsum.photos/200/300")
	t.ProductsList = append(t.ProductsList, newElem)

	c.HTML(http.StatusOK, "todo-list-elem", newElem)
}

func (t *ProductsHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range t.ProductsList {
		if todo.UUID == id {
			t.ProductsList = append(t.ProductsList[:i], t.ProductsList[i+1:]...)
		}
	}

	c.HTML(http.StatusOK, "todo", gin.H{
		"Todos": t.ProductsList,
	})
}

func (t *ProductsHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var elem models.Product
	for i, todo := range t.ProductsList {
		if todo.UUID == id {
			if todo.Status < 3 {
				t.ProductsList[i].Status += 1
			}
			elem = t.ProductsList[i]
		}
	}

	c.HTML(http.StatusOK, "todo-list-elem", elem)
}
