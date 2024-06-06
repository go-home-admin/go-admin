package web

import (
	"fmt"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// SwaggerUI @Bean
type SwaggerUI struct {
	http *servers.Http `inject:""`
}

func (s *SwaggerUI) Boot() {
	if app.IsDebug() {
		s.http.Static("/swagger", "./web/swagger")

		fmt.Printf("\n打开文档 http://127.0.0.1:%v/swagger/\n", s.http.Port)
	}
}
