package router

import (
	"apiserver-gin/internal/handler/ping"
	"apiserver-gin/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load 加载中间件和路由信息
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 注册中间件
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache())
	g.Use(middleware.Options())
	g.Use(middleware.Secure())
	g.Use(middleware.RequestId())
	g.Use(middleware.Logger)
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})

	// ping server
	g.GET("/ping", ping.Ping())

	// ---- user login use jwt ---
	ginJWTMiddleware := middleware.Jwt()
	g.POST("/login", ginJWTMiddleware.LoginHandler)

	// User route
	auth := g.Group("/auth")
	auth.GET("/refresh_token", ginJWTMiddleware.RefreshHandler)
	return g
}
