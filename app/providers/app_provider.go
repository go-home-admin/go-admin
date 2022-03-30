package providers

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
)

// App 系统引导结构体
// @Bean
type App struct {
	// 框架引导
	container *services.Container          `inject:""`
	frame     *providers.FrameworkProvider `inject:""`

	// 你的应用引导
	resp *Response `inject:""`
}

func (a *App) Run(servers []constraint.KernelServer) {
	a.container.Run(servers)
}
