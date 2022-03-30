package queue

import (
	"gitee.com/ctfang/go-admin/app/queue/job"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services/broker"
	log "github.com/sirupsen/logrus"
)

// @Bean
type Kernel struct {
	b      *broker.RedisBroker `inject:""`
	worker *Worker             `inject:""`
}

func (k *Kernel) Init() {
	// 工人封装
	k.b.SetWorker(k.worker)
	// 注入信息通道代理商

	// 注册Job
	k.b.Consumer(job.NewDemoJob())
}

func (k *Kernel) Run() {
	k.b.Loop()
}

func (k *Kernel) Exit() {
	log.Info("queue server exit 0")
}

func GetServer() constraint.KernelServer {
	return NewKernel()
}
