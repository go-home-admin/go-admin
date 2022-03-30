// gen for home toolset
package routes

import (
	admin_user "gitee.com/ctfang/go-admin/app/http/admin/admin_user"
	gin "github.com/gin-gonic/gin"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminRoutes struct {
	admin_user *admin_user.Controller `inject:""`
}

func (c *AdminRoutes) GetGroup() string {
	return "admin"
}
func (c *AdminRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Get("/info"): c.admin_user.GinHandleInfo,
	}
}
