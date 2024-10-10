package response

import "net/http"

// R 通用返回对象
type R struct {
	// 状态码
	Code int `json:"code"`
	// 成功
	Success bool `json:"success"`
	// 数据
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

// Success 通用成功返回
func Success() R {
	return R{
		Code:    http.StatusOK,
		Success: true,
		Data:    "",
		Msg:     "操作成功",
	}
}

// Data 成功返回数据
func Data(date interface{}) R {
	return R{
		Code:    http.StatusOK,
		Success: true,
		Data:    date,
		Msg:     "操作成功",
	}
}

// DataByCustom 成功返回数据和消息
func DataByCustom(date interface{}, msg string) R {
	return R{
		Code:    http.StatusOK,
		Success: true,
		Data:    date,
		Msg:     msg,
	}
}

// Fail 通用错误返回
func Fail() R {
	return R{
		Code:    http.StatusBadRequest,
		Success: false,
		Data:    "",
		Msg:     "操作失败",
	}
}

// FailData  通用错误返回,带错误返回描述
func FailData(date interface{}, msg string) R {
	return R{
		Code:    http.StatusInternalServerError,
		Success: false,
		Data:    date,
		Msg:     msg,
	}
}

// FailByParam 参数校验异常
func FailByParam() R {
	return R{
		Code:    http.StatusBadRequest,
		Success: false,
		Data:    "",
		Msg:     "请求参数校验失败",
	}
}

// FailByMsg 返回错误消息
func FailByMsg(msg string) R {
	return R{
		Code:    http.StatusBadRequest,
		Success: false,
		Data:    "",
		Msg:     msg,
	}
}

func FailByMsgError(msg string) R {
	return R{
		Code:    http.StatusInternalServerError,
		Success: false,
		Data:    "",
		Msg:     msg,
	}
}

// FailByCustom 自定义错误返回
func FailByCustom(code int, msg string) R {
	return R{
		Code:    code,
		Success: false,
		Data:    "",
		Msg:     msg,
	}
}
