package job

import (
	"gitee.com/ctfang/go-admin/app/message/event"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/utils"
	"time"
)

// @Bean
type DemoJob struct {
	Test string
}

func (d *DemoJob) Config() constraint.JobConfig {
	return constraint.JobConfig{
		Event: event.SmsTask{},
	}
}

func (d *DemoJob) Handler(task constraint.Task) {
	utils.Dump(d.Test)
	utils.Dump(111)

	d.Test = time.Now().String()
}
