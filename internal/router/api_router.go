// Created on 2023/3/4.
// @author tony
// email xmgtony@gmail.com
// description

package router

import (
	"apiserver-gin/internal/handler/v1/accountbill"
	"apiserver-gin/internal/handler/v1/user"
	"apiserver-gin/internal/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	userHandler        *user.Handler
	accountBillHandler *accountbill.Handler
}

func NewApiRouter(
	userHandler *user.Handler,
	accountBillHandler *accountbill.Handler) *ApiRouter {
	return &ApiRouter{
		userHandler:        userHandler,
		accountBillHandler: accountBillHandler,
	}
}

// Load 实现了server/http.go:40
func (ar *ApiRouter) Load(g *gin.Engine) {
	// login
	g.POST("/login", ar.userHandler.Login())
	// user group
	ug := g.Group("/v1/user", middleware.AuthToken())
	{
		ug.GET("", ar.userHandler.GetUserInfo())
	}

	// accountBill group
	abg := g.Group("/v1/accountBill", middleware.AuthToken())
	{
		abg.GET("/list", ar.accountBillHandler.GetAccountBillList())
		// login
		abg.POST("", ar.accountBillHandler.AddAccountBill())
	}
}
