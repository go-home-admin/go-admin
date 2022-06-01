// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	public "github.com/go-home-admin/go-admin/app/http/admin/public"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminPublicRoutes struct {
	public *public.Controller `inject:""`
}

func (c *AdminPublicRoutes) GetGroup() string {
	return "admin-public"
}
func (c *AdminPublicRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Post("/logout"): c.public.GinHandleLogout,
		api.Post("/login"):  c.public.GinHandleLogin,
	}
}
