package tools

import "github.com/satori/go.uuid"

// GenUUID 生成一个随机的唯一ID
func GenUUID() string {
	return uuid.NewV4().String()
}

// GenUUIDFromStr 从指定的字符串生成uuid
func GenUUIDFromStr(str string) (string, error) {
	u, err := uuid.FromString(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
