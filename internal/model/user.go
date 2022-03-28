package model

import (
	"apiserver-gin/pkg/time"
)

// User 对应数据库user表
type User struct {
	BaseModel
	Name     string        `gorm:"column:name" json:"name"`
	Password string        `gorm:"column:password" json:"-"` // 密码json化时要忽略避免泄露，用不到时sql中不要查询该字段
	Mobile   string        `gorm:"column:mobile" json:"mobile"`
	Email    string        `gorm:"column:email" json:"email"`
	Sex      uint          `gorm:"column:sex" json:"sex"`
	Age      uint          `gorm:"column:age" json:"age"`
	Birthday time.JsonTime `gorm:"column:birthday" json:"birthday"`
}

func (User) TableName() string {
	return "user"
}

//LoginReq 登录请求，登录标识ID需要为邮件或者手机号码，密码介于6-32之间
type LoginReq struct {
	ID       string `json:"ID" validate:"required,email|mobile"`
	Password string `json:"password" binding:"required,gte=6,lte=32"`
}
