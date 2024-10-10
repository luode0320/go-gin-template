// controllers/user_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/common/response"
	"go-gin-template/db"
	"go-gin-template/services"
	"net/http"
)

type UserController struct{}

var service = services.UserService{}

func (uc *UserController) GetUsers(c *gin.Context) {
	users := service.GetUsers()
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := service.GetUser(id)
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	user := c.MustGet("form").(*db.User)
	service.CreateUser(user)
	c.JSON(http.StatusCreated, response.Data(user))
}
