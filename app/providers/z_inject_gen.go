// gen for home toolset
package providers

import (
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
		_AppSingle.MysqlProvider = providers_0.NewMysqlProvider()
		_AppSingle.RedisProvider = providers_0.NewRedisProvider()
		_AppSingle.Response = NewResponse()
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
