package db

import (
	"database/sql"
	"github.com/Alfagov/Todo_HTMX_MVC/models"
)

type Dao interface {
	CreateUser(user *models.User) error
	LoginUser(username string, password string) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	ExistsUserByUsername(username string) (bool, error)
	ExistsUserByEmail(email string) (bool, error)

	CreateTodoList(todoList *models.TodoList) error
	GetTodoListByName(name string, userId string) (*models.TodoList, error)
	GetTodoListOfUser(userId string) ([]models.TodoList, error)
	DeleteTodoListByName(name string, userId string) error
}

type DaoImpl struct {
	*sql.DB
}

func NewDao(db *sql.DB) Dao {
	return &DaoImpl{db}
}

func (d *DaoImpl) ExistsUserByEmail(email string) (bool, error) {
	err := d.QueryRow("SELECT email FROM users WHERE email = ?;", email).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *DaoImpl) LoginUser(username string, password string) (*models.User, error) {
	var user models.User
	err := d.QueryRow("SELECT * FROM users WHERE username = ? AND password = ?;", username, password).Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (d *DaoImpl) ExistsUserByUsername(username string) (bool, error) {
	err := d.QueryRow("SELECT username FROM users WHERE username = ?;", username).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *DaoImpl) CreateUser(user *models.User) error {
	_, err := d.Exec("INSERT INTO users (id, username, name, email, password) VALUES (?, ?, ?, ?, ?);",
		user.Id, user.Username, user.Name, user.Email, user.Password)
	return err
}

func (d *DaoImpl) GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := d.QueryRow("SELECT * FROM users WHERE name = ?;", name).Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password)
	return &user, err
}

func (d *DaoImpl) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := d.QueryRow("SELECT * FROM users WHERE username = ?;", username).Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password)
	return &user, err
}

func (d *DaoImpl) CreateTodoList(todoList *models.TodoList) error {
	_, err := d.Exec("INSERT INTO todos (id, name, desc, status, user_id) VALUES (?, ?, ?, ?, ?);",
		todoList.Id, todoList.Name, todoList.Desc, todoList.Status, todoList.UserId)
	return err
}

func (d *DaoImpl) GetTodoListByName(name string, userId string) (*models.TodoList, error) {
	var todoList models.TodoList
	err := d.QueryRow("SELECT * FROM todos WHERE name = ? AND user_id = ?;", name, userId).Scan(&todoList.Id, &todoList.Name, &todoList.Desc, &todoList.Status, &todoList.UserId)
	return &todoList, err
}

func (d *DaoImpl) queryTodosByUserId(userId string) (dbTodoList, error) {
	rows, err := d.Query("SELECT * FROM todos WHERE user_id = ?;", userId)
	if err != nil {
		return dbTodoList{}, err
	}
	return dbTodoList{rows}, nil
}

func (d *DaoImpl) GetTodoListOfUser(userId string) ([]models.TodoList, error) {

	rows, err := d.queryTodosByUserId(userId)
	if err != nil {
		return nil, err
	}
	todoList, err := rows.ScanAll()
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

func (d *DaoImpl) DeleteTodoListByName(name string, userId string) error {
	_, err := d.Exec("DELETE FROM todos WHERE name = ? AND user_id = ?;", name, userId)
	return err
}
