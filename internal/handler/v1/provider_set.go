// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description user handler层ProviderSet

package user

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserHandler)
