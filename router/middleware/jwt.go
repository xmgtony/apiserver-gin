package middleware

import (
	usermodel "apiserver-gin/model/user"
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/errcode"
	. "apiserver-gin/pkg/log"
	"apiserver-gin/pkg/response"
	time2 "apiserver-gin/pkg/time"
	"apiserver-gin/tools/security"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

var identityKey string = "jwt-key"

func Jwt() *jwt.GinJWTMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                 config.Cfg.ApplicationName,
		Key:                   []byte(config.Cfg.JwtSecret),
		Timeout:               time.Hour * 24,
		MaxRefresh:            time.Hour * 24,
		IdentityKey:           identityKey,
		TokenLookup:           "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:         "Bearer",
		PayloadFunc:           payloadFunc,
		IdentityHandler:       identityHandler,
		Authenticator:         authenticator,
		Authorizator:          authorizator,
		Unauthorized:          unauthorized,
		LoginResponse:         loginResponse,
		HTTPStatusMessageFunc: HTTPStatusMessageFunc,
		TimeFunc:              time.Now,
		SendCookie:            false,
		SecureCookie:          false,
		CookieHTTPOnly:        true,
		SendAuthorization:     false,
		DisabledAbort:         false,
	})
	if err != nil {
		Log.Fatal("Jwt create error", zap.String("error", err.Error()))
	}
	return jwtMiddleware
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*usermodel.User); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	if v, ok := claims[identityKey].(*usermodel.User); ok {
		return v
	}
	return nil
}

func authenticator(c *gin.Context) (interface{}, error) {
	var userLogin usermodel.UserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	user := &usermodel.User{
		Name:     userLogin.Name,
		Password: userLogin.Password,
	}
	err := user.Validate()
	if err != nil {
		return nil, errcode.ValidateErr
	}
	password := strings.Trim(userLogin.Password, " ")
	// 查询用户相关信息，检验用户名密码是否正确。
	user, err = usermodel.GetUser(userLogin.Name)
	if err == nil && security.ValidatePassword(password, user.Password) {
		return user, nil
	}
	return nil, errcode.UserLoginErr
}

func authorizator(data interface{}, c *gin.Context) bool {
	// 暂时不会对用户做相关模块权限控制，默认登录后就可以访问整个系统，所以始终返回true
	// 请根据自己业务来按需编写授权规则
	return true
}

func unauthorized(c *gin.Context, code int, message string) {
	response.SendJson(c, errcode.NewCode(code, message), nil)
}

// loginResponse 自定义响应格式，与整个应用统一
func loginResponse(c *gin.Context, code int, jwtToken string, expire time.Time) {
	respMap := make(map[string]interface{})
	respMap["token"] = jwtToken
	respMap["expire"] = time2.JsonTime(expire)
	response.SendJson(c, nil, respMap)
}

func HTTPStatusMessageFunc(e error, c *gin.Context) string {
	if err, ok := e.(*errcode.Code); ok {
		return err.Message
	}
	// 需要授权
	return errcode.RequireAuth.Message
}
