package queue

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	log "github.com/sirupsen/logrus"
)

// @Bean
type Worker struct {

	// 最多并发执行
	limit     uint
	limitChan chan bool
}

func (w *Worker) Init() {
	w.limit = 100
	w.limitChan = make(chan bool, w.limit)
}

// 工作任务限制
func (w *Worker) Run(job constraint.Job, task constraint.Task) {
	go w.call(job, task)

	w.limitChan <- true
}

func (w *Worker) call(job constraint.Job, task constraint.Task) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("队列任务执行发生错误", err)
		}

		<-w.limitChan
	}()

	job.Handler(task)

}
