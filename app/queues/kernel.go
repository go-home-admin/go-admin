package queues

import (
	"github.com/go-home-admin/go-admin/app/queues/jobs"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// Kernel @Bean
type Kernel struct {
	*servers.Queue `inject:""`
}

func (k *Kernel) Init() {
	k.Listen(jobs.GetAllProvider())
	// 打开广播进程
	k.StartBroadcast()
}

func GetServer() constraint.KernelServer {
	return NewKernel()
}
