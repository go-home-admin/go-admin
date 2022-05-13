package public

import (
	admin "gitee.com/ctfang/go-admin/generate/proto/admin"
	gin "github.com/gin-gonic/gin"
	http "github.com/go-home-admin/home/app/http"
)

// Login
func (receiver *Controller) Login(req *admin.LoginRequest, ctx http.Context) (*admin.LoginResponse, error) {
	// TODO 这里写业务
	return &admin.LoginResponse{
		Token: "SCUI.Administrator.Auth",
		UserInfo: &admin.Userinfo{
			UserId:    "1",
			UserName:  "Administrator",
			Dashboard: "0",
			Role: []string{
				"SA",
				"admin",
				"Auditor",
			},
		},
	}, nil
}

// GinHandleLogin gin原始路由处理
// http.Post(/login)
func (receiver *Controller) GinHandleLogin(ctx *gin.Context) {
	req := &admin.LoginRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Login(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
