package uuid

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// GenUUID 生成一个随机的唯一ID
func GenUUID() string {
	return uuid.NewV4().String()
}

// GenUUID16 截取uuid前16位
func GenUUID16() string {
	uuidStr := uuid.NewV4().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	return uuidStr[0:16]
}

// ParseUUIDFromStr 从指定的字符串生成uuid
func ParseUUIDFromStr(str string) (string, error) {
	u, err := uuid.FromString(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
