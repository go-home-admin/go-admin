package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	"github.com/sirupsen/logrus"
	http2 "net/http"
	"strings"
)

// Response @Bean
// 重载框架的输出
type Response struct{}

func (r *Response) Init() {
	// NewContext 最好在中间件已经赋值以下两个参数
	// ginCtx.Set("user", nil)
	// ginCtx.Set("user_id", nil)
	http.NewContext = func(ctx *gin.Context) http.Context {
		return &Ctx{
			Context: ctx,
		}
	}
}

type Ctx struct {
	*gin.Context

	UserInfo interface{}
}

func (receiver Ctx) Success(data interface{}) {
	receiver.JSON(http2.StatusOK, map[string]interface{}{
		"data": data,
		"code": 0,
		"msg":  "",
	})
}

func (receiver Ctx) Fail(err error) {
	receiver.JSON(http2.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  err.Error(),
	})
}

func (receiver Ctx) Gin() *gin.Context {
	return receiver.Context
}

func (receiver Ctx) User() interface{} {
	return receiver.UserInfo
}

func (receiver Ctx) Id() uint64 {
	u, ok := receiver.Context.Get(http.UserIdKey)
	if !ok {
		logrus.Fatal("id 不存在, todo Context.Set(UserIdKey, Uid)")
		return 0
	}
	return u.(uint64)
}

func (receiver Ctx) IdStr() string {
	u, ok := receiver.Context.Get(http.UserIdKey)
	if !ok {
		return ""
	}
	return u.(string)
}

func (receiver Ctx) Token() string {
	tokenString := receiver.Context.GetHeader("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	return tokenString
}
