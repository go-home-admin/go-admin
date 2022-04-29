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
		Cors(),
	}

	// 分组中间件
	k.MiddlewareGroup = map[string][]gin.HandlerFunc{
		"admin": {},
		"api":   {},
	}
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
