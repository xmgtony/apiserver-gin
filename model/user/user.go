package user

import (
	"apidemo-gin/model"
	"errors"
	validator "gopkg.in/go-playground/validator.v9"
	"log"
)

//User 在这里是一个充血模型，即具有行为，又具有状态，一般的java中PO是贫血模型，无具体行为
type User struct {
	model.BaseModel
	Name     string `gorm:"column:name;type:varchar(32);not null" json:"name" binding:"required" validate:"min=1,max=32"`
	Password string `gorm:"column:password;type:char(32);not null" json:"-"` // 密码json化时要忽略避免泄露，用不到时sql中不要查询该字段
}

func init() {
	log.Println("go test ")
}

func (User) TableName() string {
	return "user"
}

// GetUser query user from db by name
func GetUser(name string) (*User, error) {
	if len(name) <= 0 {
		return nil, errors.New("姓名长度不能为空")
	}
	user := &User{}
	d := model.DB.Master.Where("name = ?", name).First(user)
	return user, d.Error
}

func (user *User) validate() error {
	validate := validator.New()
	return validate.Struct(user)
}
