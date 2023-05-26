package v2

import (
	"application/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) { // 接受一个 gin.Context 类型的参数 c，表示 HTTP 请求上下文
	appG := app.Gin{C: c} // 创建了一个 app.Gin 对象 appG，并将 HTTP 请求上下文 c 传递给它，以便在后续的操作中使用
	// 将 HTTP 响应以 JSON 格式返回给客户端
	appG.Response(http.StatusOK, "成功", map[string]any{
		"msg": "Hello",
	})
}
