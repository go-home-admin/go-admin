package api_demo

import (
	api "gitee.com/ctfang/go-admin/generate/proto/api"
	gin "github.com/gin-gonic/gin"
	providers "github.com/go-home-admin/home/app/providers"
)

// Home
func (receiver *Controller) Home(req *api.ApiDemoHomeRequest, ctx *gin.Context) (*api.ApiDemoHomeResponse, error) {
	// TODO 这里写业务
	return &api.ApiDemoHomeResponse{}, nil
}

// GinHandleHome gin原始路由处理
// http.Get(/api/demo)
func (receiver *Controller) GinHandleHome(ctx *gin.Context) {
	req := &api.ApiDemoHomeRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Home(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
