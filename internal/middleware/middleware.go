// author: maxf
// date: 2022-04-01 16:38
// version: 中间件初始化

package middleware

import (
	"apiserver-gin/internal/handler/ping"
	"github.com/gin-gonic/gin"
	"net/http"
)

// middleware 实现Router接口
// 便于服务启动时加载, middleware本质跟handler无区别
type middleware struct {
}

func NewMiddleware() *middleware {
	return &middleware{}
}

// Load 注册中间件和公共路由
func (m *middleware) Load(g *gin.Engine) {
	// 注册中间件
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(NoCache())
	g.Use(Options())
	g.Use(Secure())
	g.Use(RequestId())
	g.Use(Logger)
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})
	// ping server
	g.GET("/ping", ping.Ping())
}
