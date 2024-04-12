// gen for home toolset
package http

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	servers "github.com/go-home-admin/home/bootstrap/servers"
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
		_KernelSingle.Http = servers.NewHttp()
		_KernelSingle.RouteProvider = providers.NewRouteProvider()
		providers.AfterProvider(_KernelSingle, "")
	}
	return _KernelSingle
}
