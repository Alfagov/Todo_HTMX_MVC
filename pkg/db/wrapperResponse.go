package db

import (
	"database/sql"
	"github.com/Alfagov/Todo_HTMX_MVC/models"
)

type dbTodoList struct {
	*sql.Rows
}

func (tl *dbTodoList) ScanAll() ([]models.TodoList, error) {
	var todoLists []models.TodoList
	for tl.Next() {
		var todoList models.TodoList
		err := tl.Scan(&todoList.Id, &todoList.Name, &todoList.Desc, &todoList.Status, &todoList.UserId)
		if err != nil {
			return nil, err
		}
		todoLists = append(todoLists, todoList)
	}
	return todoLists, nil
}
