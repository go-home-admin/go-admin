// gen for home toolset
package providers

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _AppSingle *App
var _ResponseSingle *Response

func GetAllProvider() []interface{} {
	return []interface{}{
		NewApp(),
		NewResponse(),
	}
}

func NewApp() *App {
	if _AppSingle == nil {
		_AppSingle = &App{}
		_AppSingle.container = services.NewContainer()
		_AppSingle.frame = providers.NewFrameworkProvider()
		_AppSingle.resp = NewResponse()
		app.AfterProvider(_AppSingle, "")
	}
	return _AppSingle
}
func NewResponse() *Response {
	if _ResponseSingle == nil {
		_ResponseSingle = &Response{}
		app.AfterProvider(_ResponseSingle, "")
	}
	return _ResponseSingle
}
