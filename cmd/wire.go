// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 使用Google依赖注入工具wire

//go:build wireinject
// +build wireinject

package main

import (
	handlerV1 "apiserver-gin/internal/handler/v1"
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
	// 加载中间件及路由, 注意需要先加载中间件
	rts = append(rts,
		middleware.NewMiddleware(),
		buildUserRouter(ds),
		buildAccountBillRouter(ds),
	)
	return rts
}

var providerSet = wire.NewSet(mysql.ProviderSet, service.ProviderSet, handlerV1.ProviderSet)

// buildUserRouter 组装UserRouter
func buildUserRouter(ds db.IDataSource) server.Router {
	wire.Build(
		providerSet,
		router.UserRouterProviderSet,
	)
	return nil
}

// buildAccountBillRouter 组装AccountBillRouter
func buildAccountBillRouter(ds db.IDataSource) server.Router {
	wire.Build(
		providerSet,
		router.AccountBillRouterProviderSet,
	)
	return nil
}
