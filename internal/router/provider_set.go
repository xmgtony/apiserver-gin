// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description

package router

import (
	"apiserver-gin/server"
	"github.com/google/wire"
)

var UserRouterProviderSet = wire.NewSet(
	NewUserRouter,
	wire.Bind(new(server.Router), new(*userRouter)),
)

var AccountBillRouterProviderSet = wire.NewSet(
	NewAccountBillRouter,
	wire.Bind(new(server.Router), new(*accountBillRouter)),
)
