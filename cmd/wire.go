// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 使用Google依赖注入工具wire

//go:build wireinject
// +build wireinject

package main

import (
	"apiserver-gin/internal/handler/v1/user"
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
	rt := initRouter(ds)

	if rt != nil {
		rts = append(rts, rt)
	}
	return rts
}

func initRouter(ds db.IDataSource) server.Router {
	wire.Build(mysql.ProviderSet, service.ProviderSet, user.ProviderSet, router.ProviderSet)
	return nil
}
