// gen for home toolset
package web

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	servers "github.com/go-home-admin/home/bootstrap/servers"
)

var _AdminUISingle *AdminUI
var _SwaggerUISingle *SwaggerUI

func GetAllProvider() []interface{} {
	return []interface{}{
		NewAdminUI(),
		NewSwaggerUI(),
	}
}

func NewAdminUI() *AdminUI {
	if _AdminUISingle == nil {
		_AdminUISingle = &AdminUI{}
		_AdminUISingle.http = servers.NewHttp()
		providers.AfterProvider(_AdminUISingle, "")
	}
	return _AdminUISingle
}
func NewSwaggerUI() *SwaggerUI {
	if _SwaggerUISingle == nil {
		_SwaggerUISingle = &SwaggerUI{}
		_SwaggerUISingle.http = servers.NewHttp()
		providers.AfterProvider(_SwaggerUISingle, "")
	}
	return _SwaggerUISingle
}
