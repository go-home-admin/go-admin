// gen for home toolset
package job

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _DemoJobSingle *DemoJob

func GetAllProvider() []interface{} {
	return []interface{}{
		NewDemoJob(),
	}
}

func NewDemoJob() *DemoJob {
	if _DemoJobSingle == nil {
		_DemoJobSingle = &DemoJob{}
		app.AfterProvider(_DemoJobSingle, "")
	}
	return _DemoJobSingle
}
