// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description

package router

import (
	"apiserver-gin/server"
	"github.com/google/wire"
)

var ApiRouterProviderSet = wire.NewSet(
	NewApiRouter,
	wire.Bind(new(server.Router), new(*ApiRouter)),
)
