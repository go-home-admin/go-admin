package jobs

import (
	"gitee.com/ctfang/go-admin/app/message"
)

// DemoJob @Bean
type DemoJob struct {
	message.DemoMessage
}

func (d *DemoJob) Handler() {

}
