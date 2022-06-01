package api_demo

import (
	gin "github.com/gin-gonic/gin"
	api "github.com/go-home-admin/go-admin/generate/proto/api"
	http "github.com/go-home-admin/home/app/http"
)

// Home
func (receiver *Controller) Home(req *api.ApiDemoHomeRequest, ctx http.Context) (*api.ApiDemoHomeResponse, error) {
	// TODO 这里写业务
	return &api.ApiDemoHomeResponse{}, nil
}

// GinHandleHome gin原始路由处理
// http.Get(/api/demo)
func (receiver *Controller) GinHandleHome(ctx *gin.Context) {
	req := &api.ApiDemoHomeRequest{}
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
