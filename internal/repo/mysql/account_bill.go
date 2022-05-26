// Created on 2022/5/26.
// @author tony
// email xmgtony@gmail.com
// description 账目清单Repo

package mysql

import (
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/repo"
	"apiserver-gin/pkg/db"
	"context"
)

var _ repo.AccountBillRepo = (*accountBillRepo)(nil)

type accountBillRepo struct {
	ds db.IDataSource
}

func NewAccountBillRepo(_ds db.IDataSource) *accountBillRepo {
	return &accountBillRepo{
		ds: _ds,
	}
}

func (ab *accountBillRepo) Save(ctx context.Context, bill *model.AccountBill) error {
	return ab.ds.Master().Create(bill).Error
}

func (ab *accountBillRepo) SelectListByUserId(ctx context.Context, userId int64) ([]model.AccountBill, error) {
	var accountBills []model.AccountBill
	err := ab.ds.Master().Where("user_id = ?", userId).Find(&accountBills).Error
	if err != nil {
		return nil, err
	}
	return accountBills, nil
}
