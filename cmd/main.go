// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/config"
	"go-gin-template/db"
	"go-gin-template/routers"
	"log"
)

func init() {
	config.InitConfig()
	db.InitDB()
}

func main() {
	// 版本
	log.Printf("===========当前版本：%s==============", config.ServiceConfig.Version)

	// 设置环境debug,release,test
	gin.SetMode("release")

	// 设置gin路由
	router := routers.SetupRouter()

	// 启动服务器
	if err := router.Run(":" + config.ServiceConfig.Web.Port); err != nil {
		panic(err)
	}
}
