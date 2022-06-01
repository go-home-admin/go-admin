package jobs

import (
	"fmt"
	"github.com/go-home-admin/go-admin/app/message"
)

// DemoJob @Bean
type DemoJob struct {
	message.DemoMessage
}

func (d *DemoJob) Handler() {
	fmt.Println(d.Count)
	d.Count++
}
