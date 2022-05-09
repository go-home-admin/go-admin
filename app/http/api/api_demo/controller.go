package api_demo

import "fmt"

// Controller @Bean
type Controller struct {
	TestPort int `inject:"config, app.servers.http.port"`
}

func (receiver Controller) Init() {
	fmt.Println("启动时候打印下端口: ", receiver.TestPort)
}
