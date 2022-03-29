// author: maxf
// date: 2022-03-29 16:29
// version: 自定义校验器

package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 手机号码规则，以1开头的11位数字
var mobile, _ = regexp.Compile(`^1\d{10}$`)

func mobileValidator(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	return mobile.MatchString(phoneNumber)
}
