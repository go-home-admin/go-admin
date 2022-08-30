// gen for home toolset
package api_demo

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _ControllerSingle *Controller

func GetAllProvider() []interface{} {
	return []interface{}{
		NewController(),
	}
}

func NewController() *Controller {
	if _ControllerSingle == nil {
		_ControllerSingle = &Controller{}
		_ControllerSingle.TestPort = *providers.GetBean("config").(providers.Bean).GetBean("app.servers.http.port").(*int)
		providers.AfterProvider(_ControllerSingle, "")
	}
	return _ControllerSingle
}
