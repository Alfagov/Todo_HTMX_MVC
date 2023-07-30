package models

import "github.com/google/uuid"

type User struct {
	Id       string
	Username string
	Name     string
	Email    string
	Password string
}

func CreateUser(username string, name string, email string, password string) *User {
	return &User{Id: uuid.New().String(), Username: username, Name: name, Email: email, Password: password}
}

type TodoList struct {
	Id     string
	Name   string
	Desc   string
	Status int
	UserId string
}
