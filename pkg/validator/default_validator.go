// author: maxf
// date: 2022-03-29 11:21
// version: v1.0
// 默认validator实现，参考了部分gin源码，用于更好的支持中文提示

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
	v    Validator
)

type defaultValidator struct {
	trans ut.Translator
	v     *validator.Validate
}

func Init(language string) Validator {
	once.Do(func() {
		v = New(language, tagName)
	})
	return v
}

func New(language, tag string) *defaultValidator {
	d := &defaultValidator{
		v: validator.New(),
	}
	if len(tag) == 0 {
		tag = tagName
	}
	d.RegisterTranslator(language).RegisterTagNameFunc(tag)
	return d
}

func (d *defaultValidator) ValidStruct(data interface{}) error {
	err := d.v.Struct(data)
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		return ValidationsErrors{
			trans: d.trans,
			errs:  validationErrs,
		}
	}
	return err
}

func (d *defaultValidator) ValidStructCtx(ctx context.Context, data interface{}) error {
	err := d.v.StructCtx(ctx, data)
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		return ValidationsErrors{
			trans: d.trans,
			errs:  validationErrs,
		}
	}
	return err
}

func (d *defaultValidator) RegisterTranslator(language string) Validator {
	zhTrans := zh.New()
	enTrans := en.New()
	uni := ut.New(zhTrans, enTrans)
	trans, _ := uni.GetTranslator(language)
	// 不用考虑为空情况，不符合zh，en时会默认返回第一个fallback，即zhTrans
	switch language {
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(d.v, trans)
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(d.v, trans)
	}
	d.trans = trans
	return d
}

func (d *defaultValidator) RegisterTagNameFunc(tagName string) Validator {
	d.v.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(tagName)
	})
	return d
}

func (d *defaultValidator) ValidatorEngine() *validator.Validate {
	return d.v
}

// RegisterValidation 注册自定义标签校验器，并且注册校验器对应的提示信息。
// 否则只有自定义校验器，但是没有对应中文提示很不友好，翻译只能翻译预设的标签，自定义标签需要自己添加提示消息。
func (d *defaultValidator) RegisterValidation(tagName, Msg string, f RegisterFunc) error {
	if err := d.v.RegisterValidation(tagName, validator.Func(f)); err != nil {
		return err
	}
	// 注册标签对应提示信息翻译器
	if err := d.RegisterTagTranslator(tagName, Msg); err != nil {
		return err
	}
	return nil
}

func (d *defaultValidator) RegisterTagTranslator(tag string, msg string) error {
	f := func(ut ut.Translator) error {
		if err := d.trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
	return d.v.RegisterTranslation(tag, d.trans, f, translateFn)
}

func translateFn(trans ut.Translator, f validator.FieldError) string {
	msg, err := trans.T(f.Tag(), f.Field())
	if err == nil {
		return msg
	}
	return ""
}

func Engine() *validator.Validate {
	return v.ValidatorEngine()
}

func Struct(data interface{}) error {
	return v.ValidStruct(data)
}

func StructCtx(ctx context.Context, data interface{}) error {
	return v.ValidStructCtx(ctx, data)
}
