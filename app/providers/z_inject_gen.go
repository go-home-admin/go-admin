// gen for home toolset
package providers

import (
	web "github.com/go-home-admin/go-admin/web"
	providers_0 "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
)

var _AppSingle *App
var _AuthSingle *Auth
var _ResponseSingle *Response

func GetAllProvider() []interface{} {
	return []interface{}{
		NewApp(),
		NewAuth(),
		NewResponse(),
	}
}

func NewApp() *App {
	if _AppSingle == nil {
		_AppSingle = &App{}
		_AppSingle.Container = services.NewContainer()
		_AppSingle.FrameworkProvider = providers_0.NewFrameworkProvider()
		_AppSingle.SwaggerUI = *web.NewSwaggerUI()
		_AppSingle.AdminUI = *web.NewAdminUI()
		providers_0.AfterProvider(_AppSingle, "")
	}
	return _AppSingle
}
func NewAuth() *Auth {
	if _AuthSingle == nil {
		_AuthSingle = &Auth{}
		providers_0.AfterProvider(_AuthSingle, "")
	}
	return _AuthSingle
}
func NewResponse() *Response {
	if _ResponseSingle == nil {
		_ResponseSingle = &Response{}
		providers_0.AfterProvider(_ResponseSingle, "")
	}
	return _ResponseSingle
}
