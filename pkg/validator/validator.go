// author: maxf
// date: 2022-03-28 15:30
// version: 校验器

package validator

import (
	"context"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"sync"
)

const tagName = "label"

var (
	once sync.Once
	v    *validator.Validate
)

func Init(language string) {
	once.Do(func() {
		v = validator.New()
		registerTranslator(v, language)
		registerTagNameFunc(v, tagName)
	})
}

func registerTranslator(v *validator.Validate, language string) *validator.Validate {
	zhTrans := zh.New()
	enTrans := en.New()
	uni := ut.New(zhTrans, enTrans)
	trans, _ := uni.GetTranslator(language)
	// 不用考虑为空情况，不符合zh，en时会默认返回第一个fallback，即zhTrans
	switch language {
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return v
}

func registerTagNameFunc(v *validator.Validate, tagName string) {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(tagName)
	})
}

func Struct(data interface{}) error {
	return v.Struct(data)
}

func StructCtx(ctx context.Context, data interface{}) error {
	return v.StructCtx(ctx, data)
}
