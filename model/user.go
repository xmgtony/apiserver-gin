package model

//User 在这里是一个充血模型，即具有行为，又具有状态，一般的java中PO是贫血模型，无具体行为
type User struct {
	BaseModel
	Name     string `gorm:"column:name;type:varchar(32);not null" json:"name"`
	Password string `gorm:"column:password;type:char(32);not null" json:"-"` // 密码json化时要忽略避免泄露，用不到时sql中不要查询该字段
}

func (User) TableName() string {
	return "user"
}
