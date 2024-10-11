// routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/middleware/cache"
	"go-gin-template/middleware/cors"
	"go-gin-template/middleware/log"
	"go-gin-template/middleware/request_timeout"
	"go-gin-template/middleware/validator"
	"go-gin-template/mvc/controllers"
	"go-gin-template/mvc/model"
)

// SetupRouter 设置gin路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 创建日志中间件
	router.Use(log.CustomLogger())
	// 创建请求超时时间中间件
	router.Use(request_timeout.RequestTimeout())
	// 跨域 CORS 中间件
	router.Use(cors.Cors())
	// 创建验证器中间件实例
	valid := validator.NewFormValidator()

	// 版本 1 的路由
	api1 := router.Group("/api/v1.0")
	{
		// 注册路由处理器
		userCtrl1 := &controllers.UserController{}
		api1.GET("/users", userCtrl1.GetUsers)
		api1.GET("/users/:id", cache.Cache(), userCtrl1.GetUser)                              // 添加缓存中间件
		api1.POST("/users/create", valid.ValidateForm(new(model.User)), userCtrl1.CreateUser) // 添加验证json参数中间件
	}

	// 版本 2 的路由
	api2 := router.Group("/api/v2.0")
	{
		// 注册路由处理器
		userCtrl2 := &controllers.UserController{}
		api2.GET("/users", userCtrl2.GetUsers)
		api2.GET("/users/:id", cache.Cache(), userCtrl2.GetUser)                              // 添加缓存中间件
		api2.POST("/users/create", valid.ValidateForm(new(model.User)), userCtrl2.CreateUser) // 添加验证json参数中间件
	}

	return router
}
