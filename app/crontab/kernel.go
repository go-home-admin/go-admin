package crontab

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
	"github.com/sirupsen/logrus"
)

// Kernel @Bean
type Kernel struct {
	*servers.Crontab `inject:""`
}

func (k *Kernel) Init() {
	// 设置定时任务
	_, _ = k.AddFunc("* * * * * *", func() {
		logrus.Debug("crontab")
	})
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
