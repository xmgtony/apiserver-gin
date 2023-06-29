// author: xmgtony
// date: 2023-06-29 15:00
// version: 事务操作演示

package service

import (
	"apiserver-gin/pkg/db"
	"context"
)

// TxDemoService txDemo服务接口
type TxDemoService interface {
	SaveWithTx(ctx context.Context)
}

// txDemoService 默认实现
type txDemoService struct {
	userService UserService
	billService AccountBillService
	tx          db.Transaction
}

func NewTxDemoService(us UserService, bs AccountBillService, tx db.Transaction) *txDemoService {
	return &txDemoService{
		userService: us,
		billService: bs,
		tx:          tx,
	}
}

func (tds *txDemoService) SaveWithTx(ctx context.Context) {
	err := tds.tx.Execute(ctx, func(context context.Context) error {
		// TODO 这里只是举例，实际请根据业务执行多个service操作
		// 操作1
		// tds.userService.Save(context, user)
		// 操作2
		// tds.billService.Save(context, bill)
		//if (条件1) {
		//	返回err则回滚事务
		//	return err
		//}
		return nil
	})
	if err != nil {
		// 处理error
		return
	}
}
