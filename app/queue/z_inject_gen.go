// gen for home toolset
package queue

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
	broker "github.com/go-home-admin/home/bootstrap/services/broker"
)

var _KernelSingle *Kernel
var _WorkerSingle *Worker

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
		NewWorker(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		_KernelSingle = &Kernel{}
		_KernelSingle.worker = NewWorker()
		_KernelSingle.b = broker.NewRedisBroker()
		app.AfterProvider(_KernelSingle, "")
	}
	return _KernelSingle
}
func NewWorker() *Worker {
	if _WorkerSingle == nil {
		_WorkerSingle = &Worker{}
		app.AfterProvider(_WorkerSingle, "")
	}
	return _WorkerSingle
}
