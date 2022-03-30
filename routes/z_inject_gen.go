// gen for home toolset
package routes

import (
	admin_user "gitee.com/ctfang/go-admin/app/http/admin/admin_user"
	public "gitee.com/ctfang/go-admin/app/http/admin/public"
	api_demo "gitee.com/ctfang/go-admin/app/http/api/api_demo"
	public1 "gitee.com/ctfang/go-admin/app/http/api/public"
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _AdminRoutesSingle *AdminRoutes
var _AdminPublicRoutesSingle *AdminPublicRoutes
var _ApiRoutesSingle *ApiRoutes

func GetAllProvider() []interface{} {
	return []interface{}{
		NewAdminRoutes(),
		NewAdminPublicRoutes(),
		NewApiRoutes(),
	}
}

func NewAdminRoutes() *AdminRoutes {
	if _AdminRoutesSingle == nil {
		_AdminRoutesSingle = &AdminRoutes{}
		_AdminRoutesSingle.admin_user = admin_user.NewController()
		app.AfterProvider(_AdminRoutesSingle, "")
	}
	return _AdminRoutesSingle
}
func NewAdminPublicRoutes() *AdminPublicRoutes {
	if _AdminPublicRoutesSingle == nil {
		_AdminPublicRoutesSingle = &AdminPublicRoutes{}
		_AdminPublicRoutesSingle.public = public.NewController()
		app.AfterProvider(_AdminPublicRoutesSingle, "")
	}
	return _AdminPublicRoutesSingle
}
func NewApiRoutes() *ApiRoutes {
	if _ApiRoutesSingle == nil {
		_ApiRoutesSingle = &ApiRoutes{}
		_ApiRoutesSingle.public = public1.NewController()
		_ApiRoutesSingle.api_demo = api_demo.NewController()
		app.AfterProvider(_ApiRoutesSingle, "")
	}
	return _ApiRoutesSingle
}
