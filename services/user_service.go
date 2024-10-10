// services/user_service.go
package services

import (
	_ "github.com/mattn/go-sqlite3"
	"go-gin-template/db"
)

type UserService struct{}

func (uc *UserService) GetUsers() []db.User {
	var users []db.User
	db.DB.Find(&users)
	return users
}

func (uc *UserService) GetUser(id string) db.User {
	var user db.User
	db.DB.First(&user, id)
	return user
}

func (uc *UserService) CreateUser(user *db.User) {
	db.DB.Create(&user)
}
