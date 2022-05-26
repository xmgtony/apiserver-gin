// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description user handler层ProviderSet

package v1

import (
	"apiserver-gin/internal/handler/v1/accountbill"
	"apiserver-gin/internal/handler/v1/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.NewUserHandler,
	accountbill.NewAccountBillHandler,
)
