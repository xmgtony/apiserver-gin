// Created on 2021/3/12.
// @author tony
// email xmgtony@gmail.com
// description 用户服务

package service

import (
	"apiserver-gin/internal/model/user"
	"apiserver-gin/pkg/errcode"
	"context"
)

type UserService struct {
	userModel user.User
}

func NewUserService(user user.User) *UserService {
	return &UserService{userModel: user}
}

// GetByName 通过用户名/ID 查找用户
func (s *UserService) GetByName(ctx context.Context, name string) (*user.User, error) {
	if len(name) == 0 {
		return nil, errcode.UserLoginErr
	}
	return s.userModel.GetByName(ctx, name)
}
