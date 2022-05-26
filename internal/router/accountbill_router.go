// Created on 2022/5/27.
// @author tony
// email xmgtony@gmail.com
// description 账目清单router

package router

import (
	"apiserver-gin/internal/handler/v1/accountbill"
	"apiserver-gin/internal/middleware"
	"github.com/gin-gonic/gin"
)

// accountBillRouter 账目清单router
type accountBillRouter struct {
	abh *accountbill.handler
}

func NewAccountBillRouter(_abh *accountbill.handler) *accountBillRouter {
	return &accountBillRouter{
		abh: _abh,
	}
}

// Load 加载中间件和路由信息
func (abr *accountBillRouter) Load(g *gin.Engine) {
	// user group
	ug := g.Group("/v1/accountBill", middleware.AuthToken())
	{
		ug.GET("/list", abr.abh.GetAccountBillList())
		// login
		ug.POST("", abr.abh.AddAccountBill())
	}
}
