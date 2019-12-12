package error

import "fmt"

// 自定义错误, 也称业务错误。包含错误码错误信息和错误栈
type BizErr struct {
	Code
	Err error
}

func (b BizErr) Error() string {
	return b.Message
}

func New(code Code, err error) *BizErr {
	return &BizErr{
		Code: code,
		Err:  err,
	}
}
// Append 用来在错误信息上追加自己传递的message
func (b *BizErr) Append(message string) error {
	b.Message += "," + message
	return b
}
// Appendf 用来追加并格式化错误信息
func (b *BizErr) Appendf(format string, args ...interface{}) error {
	b.Message += "," + fmt.Sprintf(format, args)
	return b
}
// DecodeBizErr 用来解err，将err还原为 code和message
func DecodeBizErr(err error) (int, string)  {
	if err == nil {
		return SUCCESS.ErrCode, SUCCESS.Message
	}
	bizErr, ok := err.(BizErr)
	if !ok {
		// 返回系统错误
		return SystemErr.ErrCode, SystemErr.Message
	}
	return bizErr.ErrCode, bizErr.Message
}
