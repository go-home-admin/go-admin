package auth

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Context struct {
	gin *gin.Context
}

// 来自api的初始化
func NewContext(ctx *gin.Context) *Context {
	c := Context{gin: ctx}
	return &c
}

// 来自rpc的初始化
func NewContextFormGrpc(ctx *context.Context) *Context {
	c := Context{}
	return &c
}

func (c *Context) Gin() *gin.Context {
	return c.gin
}

// GetId 获取登陆用户ID, 这个获取不涉及IO, 直接解析token得到
func (c *Context) GetId() int32 {
	return 0
}
