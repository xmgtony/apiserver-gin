// Created on 2021/3/13.
// @author tony
// email xmgtony@gmail.com
// description 用户持久层

package dao

import (
	"apiserver-gin/internal/model/user"
	"context"
)

type UserDao struct {
	*Dao
}

func (ud *UserDao) GetByName(ctx context.Context, name string) (*user.User, error) {
	u := &user.User{}
	d := ud.db.Where("name = ?", name).First(u)
	return u, d.Error
}
