// author: maxf
// date: 2022-03-28 17:14
// version:

package validator

import (
	"testing"
)

type UserInfo struct {
	Name string `validate:"required,min=0,max=32"`
	Age  uint   `validate:"required,gte=0,lte=100"`
}

type PersonalProfile struct {
	Province string `validate:"min=0,max=120"`
	City     string `validate:"min=0,max=120"`
	Email    string `validate:"-"`
	Mobile   string `validate:"mobile" label:"手机号码"`
}

func init() {
	Init("zh")
}

func TestStruct(t *testing.T) {
	u := UserInfo{Age: 101}
	err := Struct(u)
	if err != nil {
		t.Logf("%v\r\n", err) // 打印 “Name为必填字段,Age必须小于或等于100”
	}

	u1 := UserInfo{
		Name: "xmgtony",
		Age:  20,
	}
	err = Struct(u1)
	if err != nil {
		t.Logf("%v", err)
	}
}

func TestLazyInitGinValidator(t *testing.T) {
	// 替换gin实现，并绑定系统validator默认实现
	LazyInitGinValidator("zh")

	p := PersonalProfile{
		Province: "浙江省",
		City:     "杭州市",
		Email:    "xmgtonygmail.com",
		Mobile:   "1901111111", // 10位数字不合法
	}
	err := Struct(p)
	if err == nil {
		t.Error("Expect an error")
	} else {
		t.Logf("%v", err) // 打印 "手机号码必须为11位数字"
	}
}
