# go-admin
![image](https://github.com/go-home-admin/toolset/blob/main/show.gif)
#### 介绍
实践一个高度面向业务的框架，面向业务，应该每个常用服务都是属于开箱即用的的状态，通用服务也使用统一风格提供，可插拔的组件方式封装
[查看文档](https://learnku.com/docs/go-admin)

#### 软件架构
软件架构说明


#### 安装教程

1. mac 安装工具链 make mac-install
2. 环境变量 copy .env.copy .env
3. 启动前运行命令, toolset make
4. 启动go run main.go

#### 使用说明

1.  生成文档: `toolset make:swagger`
2.  xxxx
3.  xxxx

#### 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request

## 快速入门
### 添加新的模块和api
在home/protobuf创建新的目录admin, 创建proto文件, 同时设置和目录一致的package和option
~~~~protobuf
package admin;
option go_package = "github.com/go-home-admin/home/generate/proto/admin";
service Public {
  option (http.RouteGroup) = "admin-public";

  rpc Index(IndexRequest)returns(IndexResponse){
    option (http.Get) = "/hello";
  }
}
~~~~
在这里app/http/kernel.go, 需要注册你的业务前缀, 中间件。

执行`toolset make`就会生成文档和基础代码, 这时项目已经正启动和访问, 当然访问api只响应基础字段。

### 引导和依赖注入
在任意地方声明方法集合时, 加上@Bean就会被框架自动实例
~~~~go
// @Bean
type Controller struct {
    // logic 会被自动注入, 方法可以直接使用
    logic *logic.Demo `inject:""`
}
func (receiver *Controller) Test() {
    logic.Call()
}
~~~~
如果是在测试文件时, 不在引导流程的的结构体, 也可以正常获得并调用, InitializeNew{你的结构体}Provider是工具生成的
~~~~go
func Test_Login(t *testing.T) {
    InitializeNewControllerProvider().Test()
}
~~~~

### 注册自己的服务
~~~~go
// @Bean
type Sms struct {}
~~~~
只要任何结构体的字段是Sms类型, 并且使用了inject标签, Sms就可以正常实例化和使用
~~~~
// @Bean
type SmsProviders struct {
    sms *Sms `inject:""`
}
// 如果实现了Init()函数, 还会被自动调用
func (i *SmsProviders) Init() {}
~~~~
