package jobs

import (
	"fmt"
	"gitee.com/ctfang/go-admin/app/message"
)

// DemoJob @Bean
type DemoJob struct {
	message.DemoMessage
}

func (d *DemoJob) Handler() {
	fmt.Println(d.Count)
	d.Count++
}
