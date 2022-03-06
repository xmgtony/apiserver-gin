package router

import (
	"apiserver-gin/internal/handler/ping"
	"apiserver-gin/internal/handler/v1/user"
	"apiserver-gin/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userHandler *user.UserHandler
)

func InitRouter(_userHandler *user.UserHandler) {
	userHandler = _userHandler
}

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
	// login
	g.POST("/login", userHandler.Login())
	// user group
	ug := g.Group("/v1/user", middleware.AuthToken())
	{
		ug.GET("", userHandler.GetUserInfo())
		// login
		ug.POST("/login", userHandler.Login())
	}
}
