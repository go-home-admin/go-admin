// gen for home toolset
package message

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _KernelSingle *Kernel

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		_KernelSingle = &Kernel{}
		app.AfterProvider(_KernelSingle, "")
	}
	return _KernelSingle
}
