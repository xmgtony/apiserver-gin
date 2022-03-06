package mysql

import (
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/repo"
	"apiserver-gin/pkg/db"
	"context"
)

type userRepo struct {
	ds db.IDataSource
}

func NewUserRepo(_ds db.IDataSource) repo.UserRepo {
	return &userRepo{
		ds: _ds,
	}
}

func (ur *userRepo) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{}
	err := ur.ds.Master().Where("name = ?", name).Find(user).Error
	return user, err
}

func (ur *userRepo) GetUserById(ctx context.Context, uid int64) (*model.User, error) {
	user := &model.User{}
	err := ur.ds.Master().Where("id = ?", uid).Find(user).Error
	return user, err
}
