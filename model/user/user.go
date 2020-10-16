package user

import (
	"apiserver-gin/model"
	"apiserver-gin/pkg/time"
	"errors"
	validator "gopkg.in/go-playground/validator.v9"
)

//User 在这里是一个充血模型，即具有行为，又具有状态，一般的java中PO是贫血模型，无具体行为
type User struct {
	model.BaseModel
	Name     string        `gorm:"column:name;type:varchar(32);not null" json:"name" validate:"min=1,max=32"`
	Password string        `gorm:"column:password;type:char(64);not null" json:"-" validate:"min=6,max=32"` // 密码json化时要忽略避免泄露，用不到时sql中不要查询该字段
	Birthday time.JsonTime `gorm:"column:birthday;type:datetime" json:"birthday"`
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

// AddUser insert into a new user on table
func (user *User) AddUser() error {
	return model.DB.Master.Create(user).Error
}

func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}
