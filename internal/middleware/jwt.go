// Created on 2021/3/21.
// @author tony
// email xmgtony@gmail.com
// description jwt中间件实现
// https://jwt.io/#libraries-io

package middleware

import "github.com/gin-gonic/gin"

// JwtMiddleware 校验jwt
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// JwtRefreshMiddleware 刷新jwt
func JwtRefreshMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// JwtLoginMiddleware 登录生成jwt
func JwtLoginMiddleware(f func()) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
