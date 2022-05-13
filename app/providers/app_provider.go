package providers

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
)

// App @Bean
// 系统引导结构体
// 所有的服务提供者都应该在这里注入(注册)
type App struct {
	*services.Container          `inject:""`
	*providers.FrameworkProvider `inject:""`
	*providers.MysqlProvider     `inject:""`
	*providers.RedisProvider     `inject:""`

	*Route    `inject:""`
	*Response `inject:""`
}

func (a *App) Run(servers []constraint.KernelServer) {
	a.Container.Run(servers)
}
