// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户信息handler

package user

import (
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/code"
	"apiserver-gin/pkg/response"
	"context"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户业务handler
type UserHandler struct {
	userSrv service.UserService
}

func NewUserHandler(_userSrv service.UserService) *UserHandler {
	return &UserHandler{
		userSrv: _userSrv,
	}
}

func (uh *UserHandler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetInt64(constant.UserID)
		user, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, user)
		}
	}
}
