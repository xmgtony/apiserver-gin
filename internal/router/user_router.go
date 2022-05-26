package router

import (
	"apiserver-gin/internal/handler/v1/user"
	"apiserver-gin/internal/middleware"
	"github.com/gin-gonic/gin"
)

// userRouter Router路由接口的默认实现
type userRouter struct {
	uh *user.handler
}

func NewUserRouter(_uh *user.handler) *userRouter {
	return &userRouter{
		uh: _uh,
	}
}

// Load 加载中间件和路由信息
func (r *userRouter) Load(g *gin.Engine) {
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
