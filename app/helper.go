package app

import (
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// PushQueue 投递进入队列的简短函数
func PushQueue(msg interface{}) {
	providers.GetBean("queue").(*servers.Queue).Push(msg)
}
