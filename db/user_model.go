// models/user_model.go
package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `form:"name" json:"name" validate:"required,min=3,max=50"`
}

func (User) TableName() string {
	return "users"
}
