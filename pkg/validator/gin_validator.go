// author: maxf
// date: 2022-03-28 14:44
// version: 基于github.com/go-playground/validator的校验

package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 手机号码规则，以1开头的11位数字
var mobile, _ = regexp.Compile(`^1-\\d{10}$`)

func LazyInitGinValidator(language string) {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 将gin默认的标签 “binding” 改为 “validate”
		// 全局一致，当struct复用时可以不用在写一遍标签
		v.SetTagName("validate")
		registerTranslator(v, language)
		registerTagNameFunc(v, tagName)
		_ = v.RegisterValidation("mobile", mobileValidator)
	}
}

func mobileValidator(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	return mobile.MatchString(phoneNumber)
}
