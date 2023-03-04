// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 使用Google依赖注入工具wire

//go:build wireinject
// +build wireinject

package main

import (
	handlerV1 "apiserver-gin/internal/handler/v1"
	"apiserver-gin/internal/repo/mysql"
	"apiserver-gin/internal/router"
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/db"
	"apiserver-gin/server"
	"github.com/google/wire"
)

// initRouter 初始化router
func initRouter(ds db.IDataSource) server.Router {
	wire.Build(
		providerSet,
		router.ApiRouterProviderSet,
	)
	return nil
}

var providerSet = wire.NewSet(mysql.ProviderSet, service.ProviderSet, handlerV1.ProviderSet)
