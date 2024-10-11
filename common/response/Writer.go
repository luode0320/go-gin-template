package response

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

// 实现一个type gin.ResponseWriter interface
type ResponseWriter struct {
	gin.ResponseWriter
	B *bytes.Buffer
}

// 重写Write([]byte) (int, error)
func (w ResponseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中再写一份数据
	w.B.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}
