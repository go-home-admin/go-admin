package main

import (
	"embed"
	"github.com/go-home-admin/go-admin/app/http"
	"github.com/go-home-admin/go-admin/app/providers"
	"github.com/go-home-admin/home/bootstrap/constraint"
	fp "github.com/go-home-admin/home/bootstrap/providers"
)

//go:embed config
var config embed.FS

func init() {
	fp.SetConfigDir(&config)
}

func main() {
	app := providers.NewApp()

	app.Run([]constraint.KernelServer{
		// http服务
		http.GetServer(),
		//election.GetServer(func() {
		//	logrus.Info("我是老大")
		//	logrus.Info("我才可以启动定时服务")
		//}),
		// Job消费服务
		// queues.GetServer(),
	})
}
