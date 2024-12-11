package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gin-template/common/response"
	"go-gin-template/config/log"
	"net/http"
)

// 验证实例
var validate = validator.New()

// ValidateForm 用于验证json数据
func ValidateForm(form interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBind(form); err != nil {
			log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, response.FailByMsg(err.Error()))
			ctx.Abort()
			return
		}

		if err := validate.Struct(form); err != nil {
			log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, response.FailByMsg(err.Error()))
			ctx.Abort()
			return
		}

		ctx.Set("form", form)
		ctx.Next()
	}
}
