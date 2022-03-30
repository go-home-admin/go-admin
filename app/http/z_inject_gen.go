// gen for home toolset
package http

import (
	help "github.com/go-home-admin/home/bootstrap/http/help"
	services "github.com/go-home-admin/home/bootstrap/services"
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _KernelSingle *Kernel

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		_KernelSingle = &Kernel{}
		_KernelSingle.httpServer = services.NewHttpServer()
		_KernelSingle.help = help.NewRouteHelp()
		_KernelSingle.config = app.GetBean("config").(app.Bean).GetBean("app").(*services.Config)
		app.AfterProvider(_KernelSingle, "")
	}
	return _KernelSingle
}
