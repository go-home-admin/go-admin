package event

import "gitee.com/ctfang/go-admin/app/message"

type SmsTask struct {
	Phone string `json:"phone"`
}

func (receiver SmsTask) Happen() {
	message.Push(receiver)
}
