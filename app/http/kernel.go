package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// Kernel @Bean
type Kernel struct {
	*servers.Http `inject:""`
}

func (k *Kernel) Init() {
	// 全局中间件
	k.Middleware = []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
	}

	// 分组中间件, 在路由提供者中自行设置
	k.MiddlewareGroup = map[string][]gin.HandlerFunc{
		"admin": {
			Cors(),
		},
		"api": {},
	}
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
