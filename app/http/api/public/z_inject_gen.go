// gen for home toolset
package public

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
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
		app.AfterProvider(_ControllerSingle, "")
	}
	return _ControllerSingle
}
