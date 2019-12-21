// 配置路由信息
package router

import (
	"apidemo-gin/handler/check"
	userv1 "apidemo-gin/handler/user/v1"
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

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "handler not found!")
	})

	// health check
	checkGroup := g.Group("/check", check.LocalIPCheck)
	{
		checkGroup.GET("/health", check.HealthCheck)
	}

	// User route
	userGroupV1 := g.Group("/v1/user")
	{
		userGroupV1.GET("/:name", userv1.Get)
		userGroupV1.POST("/create", userv1.Create)
	}
	return g
}
