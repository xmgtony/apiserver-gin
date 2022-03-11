package router

import (
	"apiserver-gin/internal/handler/ping"
	"apiserver-gin/internal/handler/v1/user"
	"apiserver-gin/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// router Router路由接口的默认实现
type router struct {
	uh *user.UserHandler
}

func NewRouter(_uh *user.UserHandler) *router {
	return &router{
		uh: _uh,
	}
}

// Load 加载中间件和路由信息
func (r *router) Load(g *gin.Engine) {
	// 注册中间件
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache())
	g.Use(middleware.Options())
	g.Use(middleware.Secure())
	g.Use(middleware.RequestId())
	g.Use(middleware.Logger)
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})

	// ping server
	g.GET("/ping", ping.Ping())
	// login
	g.POST("/login", r.uh.Login())
	// user group
	ug := g.Group("/v1/user", middleware.AuthToken())
	{
		ug.GET("", r.uh.GetUserInfo())
		// login
		ug.POST("/login", r.uh.Login())
	}
}
