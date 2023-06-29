// Created on 2023/3/15.
// @author tony
// email xmgtony@gmail.com
// description 事物控制接口

package mysql

import (
	"apiserver-gin/pkg/db"
	"context"
	"gorm.io/gorm"
)

type contextTxKey struct{}

// 事物默认实现
type transaction struct {
	ds db.IDataSource
}

func NewTransaction(_ds db.IDataSource) *transaction {
	return &transaction{ds: _ds}
}

func (t *transaction) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.ds.Master(ctx).Transaction(func(tx *gorm.DB) error {
		withValue := context.WithValue(ctx, contextTxKey{}, tx)
		return fn(withValue)
	})
}
