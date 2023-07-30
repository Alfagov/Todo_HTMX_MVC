package main

import (
	"database/sql"
	"fmt"
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/app"
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/db"
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/jwtHelper"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/libsql/libsql-client-go/libsql"
)

type Todo struct {
	Name   string
	Desc   string
	Status int
	Id     string
}

type PageData struct {
	IsAuthenticated bool
}

var todoList = []Todo{
	{"Todo 1", "Todo 1 Description", 2, uuid.New().String()},
	{"Todo 2", "Todo 2 Description", 0, uuid.New().String()},
	{"Todo 3", "Todo 3 Description", 2, uuid.New().String()},
}

func main() {

	var dbUrl = "http://127.0.0.1:4432"
	dbConn, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}

	r := gin.Default()

	dbDao := db.NewDao(dbConn)
	application := app.NewApp(dbDao, jwtHelper.NewJWT())
	application.SetupUserRoutes(r)

	r.LoadHTMLGlob("templates/*.html")
	r.Use(isHTMXMiddleware())

	r.GET("/", rootHandler)
	r.GET("/login/", loginPageHandler)
	r.POST("/login/user", loginUserHandler)
	r.POST("/add/", addHandler)
	r.DELETE("/delete/:id", removeHandler)
	r.POST("/update/:id", updateHandler)
	r.POST("/login/redirect", func(c *gin.Context) {
		c.Header("HX-Redirect", "/login/")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/register/redirect", func(c *gin.Context) {
		c.Header("HX-Redirect", "/register/")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/register/", registerPageHandler)
	r.POST("/register/user", registerUserHandler)

	r.GET("/static/css", func(c *gin.Context) {
		c.File("./templates/css/main_tw.css")
	})
	r.GET("/static/js", func(c *gin.Context) {
		c.File("./templates/js/htmx.org@1.9.4")
	})

	log.Fatal(r.Run(":8080"))

}

func isHTMXMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("HX-Request") != "" {
			c.Set("isHTMX", true)
		} else {
			c.Set("isHTMX", false)
		}

		c.Next()
	}
}

func rootHandler(c *gin.Context) {

	token, err := c.Cookie("Authentication")
	if err == nil && token != "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":           "Main website",
			"Todos":           todoList,
			"IsAuthenticated": true,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"Todos": todoList,
		})
	}
}

func loginPageHandler(c *gin.Context) {
	c.Header("HX-Redirect", "/login/")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login Page",
	})
}

func registerPageHandler(c *gin.Context) {
	c.Header("HX-Redirect", "/register/")
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register Page",
	})
}

func loginUserHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println(username, password)

	c.SetCookie("Authentication", "testToken", 3600, "/", "localhost", false, true)
	c.Header("HX-Redirect", "/")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsAuthenticated": true,
	})
}

func registerUserHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	log.Println(username, password, email)

	c.Header("HX-Redirect", "/login/")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login Page",
	})
}

func addHandler(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	log.Println(c.Cookie("Authentication"))

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
