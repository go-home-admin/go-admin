// gen for home toolset
package routes

import (
	api_demo "github.com/go-home-admin/go-admin/app/http/api/api_demo"
	public "github.com/go-home-admin/go-admin/app/http/api/public"
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _ApiRoutesSingle *ApiRoutes

func GetAllProvider() []interface{} {
	return []interface{}{
		NewApiRoutes(),
	}
}

func NewApiRoutes() *ApiRoutes {
	if _ApiRoutesSingle == nil {
		_ApiRoutesSingle = &ApiRoutes{}
		_ApiRoutesSingle.api_demo = api_demo.NewController()
		_ApiRoutesSingle.public = public.NewController()
		providers.AfterProvider(_ApiRoutesSingle, "")
	}
	return _ApiRoutesSingle
}
