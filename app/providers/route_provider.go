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

func (a *Route) Init() {
	a.LoadRoute(routes.GetAllProvider())

	a.Group("admin").Prefix("/admin")
}
