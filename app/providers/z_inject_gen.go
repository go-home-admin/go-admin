// gen for home toolset
package providers

import (
	providers_2 "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
	telescope "github.com/go-home-admin/telescope"
)

var _AppSingle *App
var _AuthSingle *Auth
var _ResponseSingle *Response
var _RouteSingle *Route

func GetAllProvider() []interface{} {
	return []interface{}{
		NewApp(),
		NewAuth(),
		NewResponse(),
		NewRoute(),
	}
}

func NewApp() *App {
	if _AppSingle == nil {
		_AppSingle = &App{}
		_AppSingle.Container = services.NewContainer()
		_AppSingle.FrameworkProvider = providers_2.NewFrameworkProvider()
		_AppSingle.MysqlProvider = providers_2.NewMysqlProvider()
		_AppSingle.RedisProvider = providers_2.NewRedisProvider()
		_AppSingle.Route = NewRoute()
		_AppSingle.Response = NewResponse()
		_AppSingle.Providers = telescope.NewProviders()
		providers_2.AfterProvider(_AppSingle, "")
	}
	return _AppSingle
}
func NewAuth() *Auth {
	if _AuthSingle == nil {
		_AuthSingle = &Auth{}
		providers_2.AfterProvider(_AuthSingle, "")
	}
	return _AuthSingle
}
func NewResponse() *Response {
	if _ResponseSingle == nil {
		_ResponseSingle = &Response{}
		providers_2.AfterProvider(_ResponseSingle, "")
	}
	return _ResponseSingle
}
func NewRoute() *Route {
	if _RouteSingle == nil {
		_RouteSingle = &Route{}
		_RouteSingle.RouteProvider = providers_2.NewRouteProvider()
		providers_2.AfterProvider(_RouteSingle, "")
	}
	return _RouteSingle
}
