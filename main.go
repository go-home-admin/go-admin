package main

import (
	"embed"
	"github.com/go-home-admin/go-admin/app/http"
	"github.com/go-home-admin/go-admin/app/providers"
	"github.com/go-home-admin/go-admin/app/queues"
	"github.com/go-home-admin/home/app/election"
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
		election.GetServer(),
		// http服务
		http.GetServer(),
		// Job消费服务
		queues.GetServer(),
	})
}
