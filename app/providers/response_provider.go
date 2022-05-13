package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	http2 "net/http"
	"strings"
)

// Response @Bean
// 重载框架的输出
type Response struct{}

func (r *Response) Init() {
	http.NewContext = func(ctx *gin.Context) http.Context {
		return &Ctx{
			Context: ctx,
		}
	}
}

type Ctx struct {
	*gin.Context
}

func (receiver Ctx) Success(data interface{}) {
	receiver.JSON(http2.StatusOK, map[string]interface{}{
		"data": data,
		"code": 200,
		"msg":  "",
	})
}

func (receiver Ctx) Fail(err error) {
	receiver.JSON(http2.StatusOK, map[string]interface{}{
		"code": 502,
		"msg":  err.Error(),
	})
}

func (receiver Ctx) Gin() *gin.Context {
	return receiver.Context
}

func (receiver Ctx) User() interface{} {
	return receiver.Context
}
func (receiver Ctx) Id() uint64 {
	return 0
}

func (receiver Ctx) Token() string {
	tokenString := receiver.Context.GetHeader("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	return tokenString
}
