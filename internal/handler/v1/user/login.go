// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户登录

package user

import (
	"apiserver-gin/internal/base/errcode"
	"apiserver-gin/internal/base/reply"
	"apiserver-gin/internal/model"
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/jwt"
	"apiserver-gin/pkg/xerrors"
	jtime "apiserver-gin/pkg/xtime"
	"apiserver-gin/tools/security"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func (uh *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginReqParam := model.LoginReq{}
		if err := c.ShouldBind(&loginReqParam); err != nil {
			reply.Fail(c, xerrors.WithCode(errcode.ValidateErr, err.Error()))
			return
		}
		// 查询用户信息
		user, err := uh.userSrv.GetByMobile(context.TODO(), loginReqParam.Mobile)
		if err != nil {
			reply.Fail(c, xerrors.WithCode(errcode.UserLoginErr, "登录失败，用户不存在"))
			return
		}

		if !security.ValidatePassword(loginReqParam.Password, user.Password) {
			reply.Fail(c, xerrors.WithCode(errcode.UserLoginErr, "登录失败，用户名、密码不匹配"))
			return
		}
		// 生成jwt token
		expireAt := time.Now().Add(24 * 7 * time.Hour)
		claims := jwt.BuildClaims(expireAt, user.Id)
		token, err := jwt.GenToken(claims, config.GlobalConfig.JwtSecret)
		if err != nil {
			reply.Fail(c, xerrors.WithCode(errcode.UserLoginErr, "生成用户授权token失败"))
			return
		}
		reply.Success(c, struct {
			Token    string     `json:"token"`
			ExpireAt jtime.Time `json:"expire_at"`
		}{
			Token:    token,
			ExpireAt: jtime.Time(expireAt),
		})
	}
}
