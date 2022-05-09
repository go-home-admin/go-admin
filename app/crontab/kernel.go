package crontab

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// Kernel @Bean
type Kernel struct {
	*servers.Crontab `inject:""`
}

func (k *Kernel) Init() {
	//_, _ = k.AddFunc("* * * * * *", func() {
	//	fmt.Println("ok")
	//})
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
