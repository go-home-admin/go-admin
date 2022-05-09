package message

type DemoMessage struct {
	Msg string
}

func (receiver DemoMessage) SetRoute() string {
	return "demo_route"
}
