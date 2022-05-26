// Created on 2022/5/27.
// @author tony
// email xmgtony@gmail.com
// description 账目清单Service
// 原则是 handler层做入参合规性校验，service层做业务校验，最终调用repo时的参数默认都是合法的。

package service

import (
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/repo"
	"context"
	"errors"
)

var _ AccountBillService = (*accountBillService)(nil)

// AccountBillService 账目清单Service接口
type AccountBillService interface {
	Save(ctx context.Context, bill *model.AccountBill) error
	SelectListByUserId(ctx context.Context, userId int64) ([]model.AccountBill, error)
}

type accountBillService struct {
	abr repo.AccountBillRepo
}

func NewAccountBillService(_abr repo.AccountBillRepo) *accountBillService {
	return &accountBillService{
		abr: _abr,
	}
}

func (abs *accountBillService) Save(ctx context.Context, bill *model.AccountBill) error {
	if nil == nil {
		return errors.New("需要保存的账目清单信息不能为空")
	}
	return abs.abr.Save(ctx, bill)
}

func (abs *accountBillService) SelectListByUserId(ctx context.Context, userId int64) ([]model.AccountBill, error) {
	if userId == 0 {
		return nil, errors.New("用户id不能为空")
	}
	return abs.abr.SelectListByUserId(ctx, userId)
}
