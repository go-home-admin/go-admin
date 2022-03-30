package public

import (
	admin "gitee.com/ctfang/go-admin/generate/proto/admin"
	gin "github.com/gin-gonic/gin"
	providers "github.com/go-home-admin/home/app/providers"
)

// Logout
func (receiver *Controller) Logout(req *admin.LogoutRequest, ctx *gin.Context) (*admin.LogoutResponse, error) {
	// TODO 这里写业务
	return &admin.LogoutResponse{}, nil
}

// GinHandleLogout gin原始路由处理
// http.Post(/logout)
func (receiver *Controller) GinHandleLogout(ctx *gin.Context) {
	req := &admin.LogoutRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Logout(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
