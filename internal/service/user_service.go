// Created on 2021/3/12.
// @author tony
// email xmgtony@gmail.com
// description 用户服务层

package service

import (
	"apiserver-gin/internal/model"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/code"
	"context"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// GetByName 通过用户名 查找用户
func (s *UserService) GetByName(ctx context.Context, name string) (*model.User, error) {
	if len(name) == 0 {
		return nil, errors.WithCode(code.ValidateErr, "用户名称不能为空")
	}
	um := &model.User{}
	return um.GetUserByName(ctx, name)
}

// GetById ID 查找用户
func (s *UserService) GetById(ctx context.Context, uid int64) (*model.User, error) {
	um := &model.User{}
	return um.GetUserById(ctx, uid)
}
