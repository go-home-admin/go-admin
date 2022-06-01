// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	admin_user "github.com/go-home-admin/go-admin/app/http/admin/admin_user"
	demo "github.com/go-home-admin/go-admin/app/http/admin/demo"
	system "github.com/go-home-admin/go-admin/app/http/admin/system"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminRoutes struct {
	admin_user *admin_user.Controller `inject:""`
	system     *system.Controller     `inject:""`
	demo       *demo.Controller       `inject:""`
}

func (c *AdminRoutes) GetGroup() string {
	return "admin"
}
func (c *AdminRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Get("/info"):           c.admin_user.GinHandleInfo,
		api.Get("/system/menu/my"): c.system.GinHandleMyMenu,
		api.Get("/demo/post"):      c.demo.GinHandlePost,
		api.Get("/demo/page"):      c.demo.GinHandlePage,
		api.Get("/demo/menu"):      c.demo.GinHandleMenu,
		api.Get("/demo/ver"):       c.demo.GinHandleVer,
	}
}
