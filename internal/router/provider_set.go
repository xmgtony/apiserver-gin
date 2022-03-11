// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description

package router

import (
	"apiserver-gin/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRouter,
	wire.Bind(new(server.Router), new(*router)),
)
