package api_demo

// Controller @Bean
type Controller struct {
	TestPort int `inject:"config, app.servers.http.port"`
}
