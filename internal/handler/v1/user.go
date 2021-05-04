// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户信息handler

package v1

import (
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/errcode"
	"apiserver-gin/pkg/response"
	"context"
	"github.com/gin-gonic/gin"
)

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, ok := c.Get(constant.UserID)
		if !ok {
			response.SendJson(c, errcode.RequireAuth, nil)
			return
		}
		uid, ok := v.(int64)
		if !ok {
			response.SendJson(c, errcode.ValidateErr, nil)
			return
		}
		user, err := s.Us.GetById(context.TODO(), uid)
		if err != nil {
			response.SendJson(c, nil, nil)
		} else {
			response.SendJson(c, nil, user)
		}
	}
}
