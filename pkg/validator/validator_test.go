// author: maxf
// date: 2022-03-28 17:14
// version:

package validator

import "testing"

type UserInfo struct {
	Name string `validate:"required,min=0,max=32"`
	Age  uint   `validate:"required,gte=0,lte=100"`
}

func init() {
	Init("zh")
}

func TestStruct(t *testing.T) {
	u := UserInfo{Age: 101}
	err := Struct(u)
	if err != nil {
		t.Logf("%v", err)
	}
}
