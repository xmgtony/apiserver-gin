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
	rts = append(rts, middleware.NewMiddleware()) // 加载中间件及公共路由
	ur := initUserRouter(ds)                      // 初始化用户路由
	if ur != nil {
		rts = append(rts, ur)
	}
	abr := initAccountBillRouter(ds) // 初始化账目清单路由
	if abr != nil {
		rts = append(rts, abr)
	}
	return rts
}

// wire使用的注意事项：
//1、如果没有在goland指定构建标签，goland无法识别wire.go, goland可能会自动把wire.go中导入的包自动去除，
// 如果出现 undeclared name: xxx 请检查导包是否正确
// 2、第一次使用wire需要安装wire本地二进制包，当生成wire_gen后可以直接通过go命令go:generate wire 从新根据wire.go生成
// 3、wire带来一定复杂度，且不完善，资料较少，官方热度不高，除几个大厂内部再用，目前很少项目使用，即使手写代码组装依赖觉得也不费事儿，只是维护没有wire方便。

// 如果需要最终生成一个实现server.Router的slice(我们期望的)，想把UserRouter和AccountBillRouter组装好，放到一起，
//目前无法实现，wire 不会把最终的成品自动合到一起。
// 另外需要注意不要在一个组装器（比如下面的initRouter里返回多个同样接口的成品结果，无论返回值是slice还是具体接口名）：
// https://github.com/google/wire/issues/214

//func initRouter(ds db.IDataSource) []server.Router {
//	// 如果你有其他路由，比如商品，可以继续添加在参数后面，每个项目组维护自己的router文件
//	// 类似 user_router.go  product_router.go 等，然后把wire需要的参数添加都下面
//	// 这样公共改动点（initRouter）很容易暴露出去，被代码审核人员发现
//	wire.Build(
//		mysql.ProviderSet,
//		service.ProviderSet,
//		handlerV1.ProviderSet,
//		router.UserRouterProviderSet,
//		router.AccountBillRouterProviderSet,
//	)
//	return nil
//}

// initUserRouter 组装UserRouter
func initUserRouter(ds db.IDataSource) server.Router {
	wire.Build(
		mysql.ProviderSet,
		service.ProviderSet,
		handlerV1.ProviderSet,
		router.UserRouterProviderSet,
	)
	return nil
}

// initAccountBillRouter 组装AccountBillRouter
func initAccountBillRouter(ds db.IDataSource) server.Router {
	wire.Build(
		mysql.ProviderSet,
		service.ProviderSet,
		handlerV1.ProviderSet,
		router.AccountBillRouterProviderSet,
	)
	return nil
}
