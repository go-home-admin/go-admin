package public

import (
	api "gitee.com/ctfang/go-admin/generate/proto/api"
	gin "github.com/gin-gonic/gin"
	providers "github.com/go-home-admin/home/app/providers"
)

// Home
func (receiver *Controller) Home(req *api.HomeRequest, ctx *gin.Context) (*api.HomeResponse, error) {
	// TODO 这里写业务
	return &api.HomeResponse{}, nil
}

// GinHandleHome gin原始路由处理
// http.Get(/)
func (receiver *Controller) GinHandleHome(ctx *gin.Context) {
	req := &api.HomeRequest{}
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
