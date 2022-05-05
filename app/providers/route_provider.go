package providers

import (
	"gitee.com/ctfang/go-admin/routes"
	"github.com/go-home-admin/home/bootstrap/providers"
)

// Route 路由服务
// @Bean
type Route struct {
	*providers.RouteProvider `inject:""`
}

// Init
// 设置路由的前缀+中间件
// 中间件使用http/kernel.go下的分组名称
// 如果不设置, 默认没有前缀
func (a *Route) Init() {
	a.LoadRoute(routes.GetAllProvider())

	a.Group("admin").Prefix("/admin").Middleware("admin")
}
