// services/user_service.go
package services

import (
	_ "github.com/mattn/go-sqlite3"
	"go-gin-template/config/db"
	"go-gin-template/model"
)

type UserService struct{}

func (uc *UserService) GetUsers() []model.User {
	var users []model.User
	db.Conn.Find(&users)
	return users
}

func (uc *UserService) GetUser(id string) model.User {
	var user model.User
	db.Conn.First(&user, id)
	return user
}

func (uc *UserService) CreateUser(user *model.User) {
	db.Conn.Create(&user)
}
