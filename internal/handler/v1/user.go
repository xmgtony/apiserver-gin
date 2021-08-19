// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户信息handler

package v1

import (
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/code"
	"apiserver-gin/pkg/response"
	"context"
	"github.com/gin-gonic/gin"
)

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetInt64(constant.UserID)
		user, err := s.Us.GetById(context.TODO(), uid)
		if err != nil {
			response.SendJson(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.SendJson(c, nil, user)
		}
	}
}
