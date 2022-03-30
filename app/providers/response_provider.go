package providers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	// ErrorRequest 参数解析错误响应
	ErrorRequest func(ctx *gin.Context, err error)
	// ErrorResponse 业务错误响应
	ErrorResponse func(ctx *gin.Context, err error)
	// SuccessResponse 成功响应
	SuccessResponse func(ctx *gin.Context, data interface{})
)

// Response @Bean
type Response struct{}

func (r *Response) Init() {
	ErrorRequest = func(ctx *gin.Context, err error) {
		Json(ctx, http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	}

	ErrorResponse = func(ctx *gin.Context, err error) {
		Json(ctx, http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	}

	SuccessResponse = func(ctx *gin.Context, data interface{}) {
		Json(ctx, http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    data,
		})
	}
}

// Json 唯一输出的出口
func Json(ctx *gin.Context, code int, h gin.H) {
	h["time"] = time.Now().Unix()
	str, _ := json.Marshal(h)
	ctx.Header("Content-Type", "application/json;charset=utf-8")
	ctx.String(code, string(str))
}
