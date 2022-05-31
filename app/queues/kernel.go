package queues

import (
	"gitee.com/ctfang/go-admin/app/queues/jobs"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// Kernel @Bean
type Kernel struct {
	*servers.Queue `inject:""`
}

func (k *Kernel) Init() {
	k.Listen(jobs.GetAllProvider())

	k.StartBroadcast()
}

func GetServer() constraint.KernelServer {
	return NewKernel()
}
