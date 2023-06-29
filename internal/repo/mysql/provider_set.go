// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description 在这里像外部提供wire工具使用的ProviderSet

package mysql

import (
	"apiserver-gin/internal/repo"
	"apiserver-gin/pkg/db"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTransaction,
	wire.Bind(new(db.Transaction), new(*transaction)),
	NewUserRepo,
	wire.Bind(new(repo.UserRepo), new(*userRepo)),
	NewAccountBillRepo,
	wire.Bind(new(repo.AccountBillRepo), new(*accountBillRepo)),
)
