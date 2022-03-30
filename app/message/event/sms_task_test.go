package event

import "testing"

func TestSmsTask_Happen(t *testing.T) {
	SmsTask{"123456"}.Happen()
}
