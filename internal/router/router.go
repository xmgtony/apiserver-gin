package router

import (
	"apiserver-gin/internal/handler"
	"apiserver-gin/internal/handler/ping"
	v1 "apiserver-gin/internal/handler/v1"
	"apiserver-gin/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load 加载中间件和路由信息
func Load(g *gin.Engine) {
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
	// user group
	ug := g.Group("/v1/user", middleware.AuthToken())
	ug.GET("", v1.GetUserInfo())
	// login
	g.POST("/login", handler.Login())
}
