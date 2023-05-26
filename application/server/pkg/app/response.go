package app

import (
	"github.com/gin-gonic/gin"
)

// Gin 将原始的 gin.Context 对象封装成一个可重用的对象，便于在不同的函数之间传递和操作
type Gin struct {
	C *gin.Context // 将 gin.Context 作为一个字段嵌入到 Gin 结构体中实现了对其功能的继承和扩展
}

// Response 用于统一向客户端返回 HTTP 响应
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, errMsg string, data any) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  errMsg,
		Data: data,
	})
	return
}
