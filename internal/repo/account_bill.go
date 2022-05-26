// Created on 2022/5/26.
// @author tony
// email xmgtony@gmail.com
// description 账目清单Repo接口， 在repo里定义接口
// 然后mysql和mongodb分别实现该接口，可以实现使用接口的地方无需变动，按需切换到不同存储设施。

package repo

import (
	"apiserver-gin/internal/model"
	"context"
)

// AccountBillRepo 账目清单Repo
type AccountBillRepo interface {
	// Save 保存账目清单数据
	Save(ctx context.Context, bill *model.AccountBill) error
	// SelectListByUserId 根据用户id查询用户账目清单
	SelectListByUserId(ctx context.Context, userId int64) ([]model.AccountBill, error)
}
