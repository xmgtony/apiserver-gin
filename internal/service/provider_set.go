// Created on 2022/3/11.
// @author tony
// email xmgtony@gmail.com
// description service层ProviderSet

package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserService,
	wire.Bind(new(UserService), new(*userService)),
)
