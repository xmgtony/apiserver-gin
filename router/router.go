package router

import (
	"apidemo-gin/handler/check"
	"apidemo-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 注册中间件
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.RequestId)
	g.Use(mw...)

	// 404 handle
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "handler not found!")
	})

	// health check
	checkGroup := g.Group("/check", check.LocalIPCheck)
	{
		checkGroup.GET("/health", check.HealthCheck)
		checkGroup.GET("/disk", check.DiskCheck)
		checkGroup.GET("/cpu", check.CPUCheck)
		checkGroup.GET("/mem", check.RAMCheck)
	}
	return g
}
