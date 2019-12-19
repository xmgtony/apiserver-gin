package security

import "testing"

func TestEncrypt(t *testing.T) {
	plaintext := "123456"
	ciphertext, err := Encrypt(plaintext)
	if err != nil {
		t.Errorf("Encrypt err %v \r\n", err)
	}
	t.Log(ciphertext)
}

func TestValidatePassword(t *testing.T) {
	plaintext := "123456"
	for i := 0; i < 10; i++ {
		ciphertext, err := Encrypt(plaintext)
		if err != nil {
			t.Fatal(err)
		}
		// 实测密码校验时速度并不怎么快，追求速度建议md5，md5应用广泛，加盐并不好破解。
		// 慢度就是bcrypt的一个设计理念，使破解时间变长。避免快速暴力破解。
		res := ValidatePassword(plaintext, ciphertext)
		if !res {
			t.Fatal("校验失败")
		}
	}
}
