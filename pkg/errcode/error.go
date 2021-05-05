package errcode

import "fmt"

// BizErr 自定义业务错误。包含错误码Code和错误栈
type BizErr struct {
	*Code
	Err error
}

func (b BizErr) Error() string {
	return b.Message
}

// Code 语义上代表错误码，上面的BizErr包含系统错误err，便于调试打印出错误栈
// 一般展示给前端时使用Code就行，Code中已经包含了前端展示所需要的基本信息
// 把Code单独拆出只是个人习惯了java中使用Enum定义错误码和错误信息的方式
// 按个人编码习惯可以将BizErr和Code合并，follow your heart
type Code struct {
	ErrCode int
	Message string
}

func (code *Code) Error() string {
	return fmt.Sprintf("code: %d, message: %s", code.ErrCode, code.Message)
}

func New(code *Code, err error) *BizErr {
	return &BizErr{
		Code: code,
		Err:  err,
	}
}

// NewCode 创建一个新的错误码, 一些场景下定义错误码比较繁琐，只是简单的展示错误信息
// 建议尽量使用定义好的错误码
func NewCode(code int, message string) *Code {
	return &Code{
		ErrCode: code,
		Message: message,
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
func DecodeBizErr(err error) (int, string) {
	if err == nil {
		return SUCCESS.ErrCode, SUCCESS.Message
	}
	switch errType := err.(type) {
	case *BizErr:
		return errType.ErrCode, errType.Message
	case *Code:
		return errType.ErrCode, errType.Message
	default:
		return SystemErr.ErrCode, SystemErr.Message
	}
}
