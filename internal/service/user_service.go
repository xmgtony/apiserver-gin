// Created on 2021/3/12.
// @author tony
// email xmgtony@gmail.com
// description 用户服务层

package service

import (
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/repo"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/code"
	"context"
)

var _ UserService = (*userService)(nil)

// UserService 定义用户操作服务接口
type UserService interface {
	// Deprecated: 使用GetByIdentification替代
	GetByName(ctx context.Context, name string) (*model.User, error)
	GetById(ctx context.Context, uid int64) (*model.User, error)
	GetByMobile(ctx context.Context, ID string) (*model.User, error)
}

// userService 实现UserService接口
type userService struct {
	ur repo.UserRepo
}

func NewUserService(_ur repo.UserRepo) *userService {
	return &userService{
		ur: _ur,
	}
}

// GetByName 通过用户名 查找用户
func (us *userService) GetByName(ctx context.Context, name string) (*model.User, error) {
	if len(name) == 0 {
		return nil, errors.WithCode(code.ValidateErr, "用户名称不能为空")
	}
	return us.ur.GetUserByName(ctx, name)
}

// GetById 根据用户ID查找用户
func (us *userService) GetById(ctx context.Context, uid int64) (*model.User, error) {
	return us.ur.GetUserById(ctx, uid)
}

// GetByMobile 根据用户手机号查询
func (us *userService) GetByMobile(ctx context.Context, mobile string) (*model.User, error) {
	// 认为handler层对service层入参都是合法的，除了业务上的校验，service层不校验入参合规性
	return us.ur.GetUserByMobile(ctx, mobile)
}
