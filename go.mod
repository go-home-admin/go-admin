module github.com/go-home-admin/go-admin

go 1.16

// 开发阶段, 把框架引导部分手动覆盖
// replace github.com/go-home-admin/home v0.0.0-20220405021323-019297df8b05 => ../home

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-home-admin/home v0.0.2
	github.com/golang/protobuf v1.5.2
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/sys v0.0.0-20220406163625-3f8b81556e12 // indirect
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.3 // indirect
	gorm.io/gorm v1.23.4
)
