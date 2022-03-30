package message

// Kernel 通知其他服务的事件，而无需担心响应
// @Bean
type Kernel struct {
}

func (k *Kernel) Init() {

}

func (k *Kernel) Push(event interface{}) {

}

// Push 事件使用这个推送
func Push(event interface{}) {
	NewKernel().Push(event)
}
