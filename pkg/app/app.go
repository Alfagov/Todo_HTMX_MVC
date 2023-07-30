package app

import (
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/db"
	"github.com/Alfagov/Todo_HTMX_MVC/pkg/jwtHelper"
	"github.com/gin-gonic/gin"
)

type App struct {
	Database db.Dao
	Router   *gin.Engine
	JWT      *jwtHelper.JWT
}

func NewApp(database db.Dao, jt *jwtHelper.JWT) *App {
	return &App{Database: database, JWT: jt}
}
