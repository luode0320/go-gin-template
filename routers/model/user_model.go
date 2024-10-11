// models/user_model.go
package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `form:"name" json:"name" validate:"required,min=3,max=50"  gorm:"unique;type:varchar(255)"`
	QQ   string `form:"qq" json:"qq" gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
