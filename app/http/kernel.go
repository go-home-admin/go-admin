package http

import (
	"fmt"
	"gitee.com/ctfang/go-admin/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/http/api"
	"github.com/go-home-admin/home/bootstrap/http/help"
	"github.com/go-home-admin/home/bootstrap/services"
)

// Kernel @Bean
type Kernel struct {
	httpServer *services.HttpServer `inject:""`
	help       *help.RouteHelp      `inject:""`
	config     *services.Config     `inject:"config, app"`

	all map[string]map[*api.Config]func(c *gin.Context)
}

// SetMiddleware
// 这里配置中间件、前缀等自定义逻辑
func (k *Kernel) SetMiddleware() []help.GroupConfig {
	return []help.GroupConfig{
		{Name: "api"},
		{Name: "swagger"},
		{Name: "admin"},
		{Name: "admin-public"},
	}
}

func (k *Kernel) Init() {
	k.all = make(map[string]map[*api.Config]func(c *gin.Context))
	k.httpServer.SetPort(k.config.GetInt("http", 80))
	k.httpServer.SetDebug(true)

	// 默认允许跨域
	engine := gin.New()
	engine.Use(Cors())
	k.httpServer.SetEngine(engine)

	for _, inf := range routes.GetAllProvider() {
		route, ok := inf.(constraint.Route)
		if ok {
			k.all[route.GetGroup()] = route.GetRoutes()
		}
	}
}

func (k *Kernel) Boot() {
	k.help.Load(k.httpServer.GetEngine(), k.SetMiddleware(), k.all)
}

func (k *Kernel) Run() {
	k.httpServer.RunListener()
}

func (k *Kernel) Exit() {
	fmt.Println("http server exit ok")
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
