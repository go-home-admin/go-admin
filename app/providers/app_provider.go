package providers

import (
	"github.com/go-home-admin/go-admin/web"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/sirupsen/logrus"
)

// App @Bean
// 系统引导结构体
// 所有的服务提供者都应该在这里注入(注册)
type App struct {
	*services.Container          `inject:""` // DI
	*providers.FrameworkProvider `inject:""` // 框架基础、默认检测是否配置 mysql、redis
	// *telescope.Providers         `inject:""` // 调试组件

	web.SwaggerUI `inject:""`
	web.AdminUI   `inject:""`

	// *Response `inject:""`
}

func (a *App) Init() {
	logrus.Debug("app init")
}

func (a *App) Boot() {
	logrus.Debug("执行完全部的init后, 再开始执行 Boot 方法")
}

func (a *App) Run(servers []constraint.KernelServer) {
	a.Container.Run(servers)
}

func (a *App) Exit() {
	logrus.Debug("app exit")
}
