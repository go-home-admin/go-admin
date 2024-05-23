GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# 安装代码工具(开发机器需要)
# export GOPATH=$HOME/go PATH=$PATH:$GOPATH/bin
mac-install:
	protoc --version || brew install protobuf						# mac下自动安装, win环境手动安装
	go install github.com/golang/protobuf/protoc-gen-go@latest		# proto 工具链, 生成go代码插件
	go install github.com/go-home-admin/toolset@latest

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./build/app ./

build_gcc:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ go build --ldflags="-linkmode external -extldflags '-static'" -o  ./build/app ./


# Orm自动维护
make-orm:
	toolset make:orm

# 只维护 protoc
protoc:
	toolset make:protoc

make-route:
	toolset make:route

make-swagger:
	toolset make:swagger
make-bean:
	toolset make:bean

# 生成全部
gen:protoc make-route make-bean make-swagger

# 调试启动
dev:gen
	go run main.go

