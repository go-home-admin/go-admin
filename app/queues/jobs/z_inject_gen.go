// gen for home toolset
package jobs

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
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
		providers.AfterProvider(_DemoJobSingle, "")
	}
	return _DemoJobSingle
}
