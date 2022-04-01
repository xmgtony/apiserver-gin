// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 使用Google依赖注入工具wire

//go:build wireinject
// +build wireinject

package main

import (
	"apiserver-gin/internal/handler/v1/user"
	"apiserver-gin/internal/middleware"
	"apiserver-gin/internal/repo/mysql"
	"apiserver-gin/internal/router"
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/db"
	"apiserver-gin/server"
	"github.com/google/wire"
)

// getRouters 获取所有的router
func getRouters(ds db.IDataSource) []server.Router {
	rts := make([]server.Router, 0)
	rts = append(rts, middleware.NewMiddleware()) // 加载中间件及公共路由
	urt := initRouter(ds)                         // 初始化用户路由
	if urt != nil {
		rts = append(rts, urt)
	}
	return rts
}

func initRouter(ds db.IDataSource) server.Router {
	// 如果你有其他路由，比如商品，可以继续添加在参数后面，每个项目组维护自己的router文件
	// 类似 user_router.go  product_router.go 等，然后把wire需要的参数添加都下面
	// 这样公共改动点（initRouter）很容易暴露出去，被代码审核人员发现
	wire.Build(
		mysql.ProviderSet,
		service.ProviderSet,
		user.ProviderSet,
		router.UserRouterProviderSet,
	)
	return nil
}
