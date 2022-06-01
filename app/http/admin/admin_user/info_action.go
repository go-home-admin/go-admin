package admin_user

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// Info
func (receiver *Controller) Info(req *admin.InfoRequest, ctx http.Context) (*admin.InfoResponse, error) {
	return &admin.InfoResponse{}, nil
}

// GinHandleInfo gin原始路由处理
// http.Get(/info)
func (receiver *Controller) GinHandleInfo(ctx *gin.Context) {
	req := &admin.InfoRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Info(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
