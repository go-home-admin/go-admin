package message

type DemoMessage struct {
	Msg   string
	Count int
}

func (receiver DemoMessage) SetRoute() string {
	return "demo_route"
}
