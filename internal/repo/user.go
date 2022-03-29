package repo

import (
	"apiserver-gin/internal/model"
	"context"
)

// UserRepo 用户repo接口
type UserRepo interface {
	// Deprecated: 使用GetUserByIdentification代替
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, uid int64) (*model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
}
