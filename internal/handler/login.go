// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户登录

package handler

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/errcode"
	"apiserver-gin/pkg/jwt"
	"apiserver-gin/pkg/response"
	"apiserver-gin/tools/security"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginParam struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		var param LoginParam
		if err := c.ShouldBind(&param); err != nil {
			response.SendJson(c, errcode.UserLoginErr, nil)
			return
		}
		// 查询用户信息
		user, err := s.Us.GetByName(context.TODO(), param.Username)
		if err != nil {
			response.SendJson(c, errcode.UserLoginErr, nil)
			return
		}

		if !security.ValidatePassword(param.Password, user.Password) {
			response.SendJson(c, errcode.UserLoginErr, nil)
			return
		}
		// 生成jwt token
		claims := jwt.BuildClaims(time.Now().Add(24*7*time.Hour), user.Id)
		token, err := jwt.GenToken(claims, config.GlobalConfig.JwtSecret)
		if err != nil {
			response.SendJson(c, errcode.UserLoginErr, nil)
			return
		}
		response.SendJson(c, nil, token)
	}
}
