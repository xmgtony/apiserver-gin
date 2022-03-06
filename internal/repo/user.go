package repo

import (
	"apiserver-gin/internal/model"
	"context"
)

// UserRepo 用户repo接口
type UserRepo interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, uid int64) (*model.User, error)
}
