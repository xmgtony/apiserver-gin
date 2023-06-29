// author: xmgtony
// date: 2023-06-29 14:38
// version: 事务接口

package db

import "context"

// Transaction 事物接口
type Transaction interface {
	// Execute 执行一个事务方法，func为一个需要保证事务完整性的业务方法
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}
