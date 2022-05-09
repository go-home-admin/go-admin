package jobs

import (
	"encoding/json"
	"gitee.com/ctfang/go-admin/app/message"
	"github.com/go-home-admin/home/bootstrap/utils"
	"reflect"
	"strings"
	"testing"
)

func TestDemoJob_Config(t *testing.T) {
	msg := message.DemoMessage{
		Msg: "ddd",
	}

	jsonStr, _ := json.Marshal(msg)

	job := &DemoJob{}
	json.Unmarshal(jsonStr, job)

	utils.Dump(jobToRoute(job))
}

func jobToRoute(handle interface{}) string {
	ref := reflect.TypeOf(handle)
	ref = ref.Elem()

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		ty := field.Type.String()
		if strings.Index(ty, "message.") != -1 {
			return ty
		}
	}
	panic("信息总线的对象路径必须包含 message. ")
}
