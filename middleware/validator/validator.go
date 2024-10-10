package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gin-template/common/response"
	"go-gin-template/middleware/log"
	"net/http"
)

// FormValidator 代表验证中间件
type FormValidator struct {
	validate *validator.Validate
}

// NewFormValidator 创建一个新的验证中间件实例
func NewFormValidator() *FormValidator {
	validate := validator.New()
	return &FormValidator{validate: validate}
}

// ValidateForm 用于验证json数据
func (fv *FormValidator) ValidateForm(form interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBind(form); err != nil {
			log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, response.FailByMsg(err.Error()))
			ctx.Abort()
			return
		}

		if err := fv.validate.Struct(form); err != nil {
			log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, response.FailByMsg(err.Error()))
			ctx.Abort()
			return
		}

		ctx.Set("form", form)
		ctx.Next()
	}
}
