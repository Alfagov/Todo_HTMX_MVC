package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	Name   string
	Desc   string
	Status int
	Id     string
}

var todoList = []Todo{
	{"Todo 1", "Todo 1 Description", 2, uuid.New().String()},
	{"Todo 2", "Todo 2 Description", 0, uuid.New().String()},
	{"Todo 3", "Todo 3 Description", 2, uuid.New().String()},
}

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("./templates/main.html")

	r.GET("/", rootHandler)
	r.POST("/add/", addHandler)
	r.DELETE("/delete/:id", removeHandler)
	r.POST("/update/:id", updateHandler)

	log.Fatal(r.Run(":8080"))

}

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": "Main website",
		"Todos": todoList,
	})
}

func addHandler(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	newElem := Todo{name, desc, 0, uuid.New().String()}
	todoList = append(todoList, newElem)

	c.HTML(http.StatusOK, "todo-list-elem", newElem)
}

func removeHandler(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todoList {
		if todo.Id == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}

	c.HTML(http.StatusOK, "todo", gin.H{
		"Todos": todoList,
	})
}

func updateHandler(c *gin.Context) {
	id := c.Param("id")
	var elem Todo
	for i, todo := range todoList {
		if todo.Id == id {
			if todo.Status < 3 {
				todoList[i].Status += 1
				elem = todoList[i]
			}
			elem = todoList[i]
		}
	}

	c.HTML(http.StatusOK, "todo-list-elem", elem)
}
