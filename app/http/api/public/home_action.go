package public

import (
	api "gitee.com/ctfang/go-admin/generate/proto/api"
	gin "github.com/gin-gonic/gin"
	http "github.com/go-home-admin/home/app/http"
)

// Home
func (receiver *Controller) Home(req *api.HomeRequest, ctx http.Context) (*api.HomeResponse, error) {
	// TODO 这里写业务
	return &api.HomeResponse{}, nil
}

// GinHandleHome gin原始路由处理
// http.Get(/)
func (receiver *Controller) GinHandleHome(ctx *gin.Context) {
	req := &api.HomeRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Home(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
