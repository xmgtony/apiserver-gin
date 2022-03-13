// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户登录

package user

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/code"
	"apiserver-gin/pkg/jwt"
	"apiserver-gin/pkg/response"
	jtime "apiserver-gin/pkg/time"
	"apiserver-gin/tools/security"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginParam struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		var param LoginParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "用户名和密码不能为空"), nil)
			return
		}
		// 查询用户信息
		user, err := uh.userSrv.GetByName(context.TODO(), param.Username)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "登录失败，用户不存在"), nil)
			return
		}

		if !security.ValidatePassword(param.Password, user.Password) {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "登录失败，用户名、密码不匹配"), nil)
			return
		}
		// 生成jwt token
		expireAt := time.Now().Add(24 * 7 * time.Hour)
		claims := jwt.BuildClaims(expireAt, user.Id)
		token, err := jwt.GenToken(claims, config.GlobalConfig.JwtSecret)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.UserLoginErr, "生成用户授权token失败"), nil)
			return
		}
		response.JSON(c, nil, struct {
			Token    string         `json:"token"`
			ExpireAt jtime.JsonTime `json:"expire_at"`
		}{
			Token:    token,
			ExpireAt: jtime.JsonTime(expireAt),
		})
	}
}
