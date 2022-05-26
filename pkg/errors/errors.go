package errors

import (
	"apiserver-gin/pkg/errors/ecode"
	"errors"
	"fmt"
	"strconv"
)

// bizErrWithCode 自定义业务错误。拓展自https://github.com/pkg/errors
type bizErrWithCode struct {
	code  int
	msg   string
	cause error
}

func (b *bizErrWithCode) Error() string {
	msg := strconv.Itoa(b.code) + ": " + b.msg
	if nil != b.cause {
		msg += ", " + b.cause.Error()
	}
	return msg
}

func (b *bizErrWithCode) Cause() error {
	return b.cause
}

func (b *bizErrWithCode) Unwrap() error {
	return b.cause
}

func (b *bizErrWithCode) GetMsg() string {
	msg := b.msg
	if b.cause != nil {
		msg += ", " + b.cause.Error()
	}
	return msg
}

func (b *bizErrWithCode) Is(err error) bool {
	if e, ok := err.(*bizErrWithCode); ok && e.code == b.code {
		return true
	}
	return false
}

func Wrap(err error, code int, msg string) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*bizErrWithCode); ok {
		return &bizErrWithCode{
			code:  e.code,
			msg:   msg,
			cause: err,
		}
	}
	return &bizErrWithCode{
		code:  code,
		msg:   msg,
		cause: err,
	}
}

func Wrapf(err error, code int, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*bizErrWithCode); ok {
		return &bizErrWithCode{
			code:  e.code,
			msg:   fmt.Sprintf(msg, args...),
			cause: err,
		}
	}
	return &bizErrWithCode{
		code:  code,
		msg:   fmt.Sprintf(msg, args...),
		cause: err,
	}
}

func WithCode(code int, msg string) *bizErrWithCode {
	return &bizErrWithCode{
		code: code,
		msg:  msg,
	}
}

// DecodeErr 用来解err，将err还原为 code和message
//func DecodeErr(err error) (int, string) {
//	if err == nil {
//		return SUCCESS.ErrCode, SUCCESS.Message
//	}
//	switch errType := err.(type) {
//	case *BizErr:
//		if errType.Err != nil {
//			errType.Append(errType.Err.Error())
//		}
//		return errType.ErrCode, errType.Message
//	case *Code:
//		return errType.ErrCode, errType.Message
//	default:
//		return SystemErr.ErrCode, SystemErr.Message
//	}
//}
// DecodeErr 用来解err，将err还原为 code和message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return ecode.Success, "success"
	}
	var b *bizErrWithCode
	if errors.As(err, &b) {
		return b.code, b.GetMsg()
	}
	return ecode.Unknown, err.Error()
}
