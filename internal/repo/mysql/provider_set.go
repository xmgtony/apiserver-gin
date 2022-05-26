// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 在这里像外部提供wire工具使用的ProviderSet

package mysql

import (
	"apiserver-gin/internal/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserRepo,
	wire.Bind(new(repo.UserRepo), new(*userRepo)),
	NewAccountBillRepo,
	wire.Bind(new(repo.AccountBillRepo), new(*accountBillRepo)),
)
