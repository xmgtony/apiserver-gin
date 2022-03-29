// author: maxf
// date: 2022-03-28 15:30
// version: 自定义校验器

package validator

import (
	"bytes"
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidStruct(data interface{}) error
	ValidStructCtx(ctx context.Context, data interface{}) error
	RegisterTranslator(language string) Validator
	RegisterTagNameFunc(tagName string) Validator
	ValidatorEngine() *validator.Validate
	RegisterValidation(tagName, Msg string, f RegisterFunc) error // 注册自定义标签，注册自定义标签的翻译信息
	RegisterTagTranslator(tag string, msg string) error           // 注册标签对应的翻译信息
}

// RegisterFunc 注册函数
type RegisterFunc func(fl validator.FieldLevel) bool

type ValidationsErrors struct {
	trans ut.Translator
	errs  validator.ValidationErrors
}

func (v ValidationsErrors) Error() string {
	translations := v.errs.Translate(v.trans)
	errBuf := bytes.NewBufferString("")
	for _, v := range translations {
		errBuf.WriteString(v)
		errBuf.WriteString(",")
	}
	errBuf.Truncate(errBuf.Len() - 1)
	return errBuf.String()
}
