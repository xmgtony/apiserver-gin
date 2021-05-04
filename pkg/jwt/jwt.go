// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description jwt实现

package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

const Issuer = "apiserver-gin"

// 在标准声明中加入用户id
type CustomClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func BuildClaims(exp time.Time, uid int64) *CustomClaims {
	return &CustomClaims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(exp),
			IssuedAt:  jwt.At(time.Now()),
			Issuer:    Issuer,
		},
	}
}

// GenToken 生成jwt token
func GenToken(c *CustomClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString([]byte(secretKey))
	return ss, err
}

// ParseToken 解析jwt token
func ParseToken(jwtStr, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
