package web

import (
	"fmt"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// AdminUI @Bean
type AdminUI struct {
	http *servers.Http `inject:""`
}

func (s *AdminUI) Boot() {
	// 显示 UI 页面
	fmt.Printf("打开 UI http://127.0.0.1:%v/admin/ui/\n\n", s.http.Port)
	s.http.Static("/admin/ui", "./web/ui")
}
