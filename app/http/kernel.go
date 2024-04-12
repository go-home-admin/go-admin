package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/routes"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// Kernel @Bean
type Kernel struct {
	*servers.Http            `inject:""`
	*providers.RouteProvider `inject:""`
}

func (k *Kernel) Init() {
	// 加载所有路由
	k.LoadRoute(routes.GetAllProvider())
	// 全局中间件
	k.Middleware = []gin.HandlerFunc{}

	// 分组中间件
	k.MiddlewareGroup = map[string][]gin.HandlerFunc{
		"admin": {
			Cors(),
		},
		"api": {},
	}

	// 如果需要, 为每个路由分组设置中间件组+前缀等
	k.Group("admin").Prefix("/admin").Middleware("admin")
	k.Group("api").Middleware("api")
}

func (k *Kernel) Boot() {

}

func (k *Kernel) Exit() {
	// http server 退出触发
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
