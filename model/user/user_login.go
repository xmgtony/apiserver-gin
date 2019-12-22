package user

// UserLogin 用户登录
type UserLogin struct {
	Name     string `form:"username" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
