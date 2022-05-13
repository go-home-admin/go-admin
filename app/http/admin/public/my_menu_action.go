package public

import (
	admin "gitee.com/ctfang/go-admin/generate/proto/admin"
	gin "github.com/gin-gonic/gin"
	http "github.com/go-home-admin/home/app/http"
)

// MyMenu   登陆
func (receiver *Controller) MyMenu(req *admin.MyMenuRequest, ctx http.Context) (*admin.MyMenuResponse, error) {
	// TODO 这里写业务
	return &admin.MyMenuResponse{}, nil
}

// GinHandleMyMenu gin原始路由处理
// http.Get(/system/menu/my)
func (receiver *Controller) GinHandleMyMenu(ctx *gin.Context) {
	req := &admin.MyMenuRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MyMenu(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
