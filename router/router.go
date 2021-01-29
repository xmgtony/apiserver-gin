// 配置路由信息
package router

import (
	"apiserver-gin/internal/handler/check"
	userv1 "apiserver-gin/internal/handler/user/v1"
	"apiserver-gin/internal/middleware"
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
	g.Use(middleware.Logger)
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

	// ---- user login use jwt ---
	ginJWTMiddleware := middleware.Jwt()
	g.POST("/login", ginJWTMiddleware.LoginHandler)

	// User route
	auth := g.Group("/auth")
	auth.GET("/refresh_token", ginJWTMiddleware.RefreshHandler)

	userGroupV1 := g.Group("/v1/user", ginJWTMiddleware.MiddlewareFunc())
	{
		userGroupV1.GET("/:name", userv1.Get)
		// 创建用户只是为了演示用法，并没有对数据库做用户名唯一性限制
		// 这样登录时，按用户名校验，如果数据库存在同一用户名用户可能会有问题
		userGroupV1.POST("/create", userv1.Create)
	}
	return g
}
