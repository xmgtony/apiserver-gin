// Created on 2022/5/26.
// @author tony
// email xmgtony@gmail.com
// description  账目清单model

package model

import (
	jtime "apiserver-gin/pkg/time"
	"github.com/shopspring/decimal"
	"time"
)

// AccountBill 账目清单
type AccountBill struct {
	BaseModel
	UserId         uint64    `gorm:"column:user_id" json:"user_id"`                 // 所属用户id
	BillDate       time.Time `gorm:"column:bill_date" json:"bill_date"`             // 账单日期
	OriginIncident string    `gorm:"column:origin_incident" json:"origin_incident"` // 账户产生的事由
	Amount         uint      `gorm:"column:amount" json:"amount"`                   // 账单金额（单位分）
	Relation       string    `gorm:"column:relation" json:"relation"`               // 与对方关系,如亲戚|同事|闺蜜
	ToName         string    `gorm:"column:to_name" json:"to_name"`                 // 对方姓名
	IsFollow       uint      `gorm:"column:is_follow" json:"is_follow"`             // 是否关注或者跟进，0不关注、1关注
	Remark         string    `gorm:"column:remark" json:"remark"`                   // 备注说明 	// 用户修改时间
}

func (m AccountBill) TableName() string {
	return "account_bill"
}

// ToAccountBillResp 转换数据库模型到响应
func (m AccountBill) ToAccountBillResp() AccountBillResp {
	// 使用BigDecimal做精确运算，避免丢失精度
	amountStr := decimal.NewFromInt32(int32(m.Amount)).
		Div(decimal.NewFromInt32(100)).StringFixed(2)
	return AccountBillResp{
		BillDate:       jtime.JsonTime(m.BillDate),
		OriginIncident: m.OriginIncident,
		Amount:         amountStr,
		Relation:       m.Relation,
		ToName:         m.ToName,
		IsFollow:       m.IsFollow,
		Remark:         m.Remark,
	}
}

// AddAccountBillReq 添加账单请求
type AddAccountBillReq struct {
	BillDate       jtime.JsonTime `json:"bill_date" validate:"required" label:"账单日期"`             // 账单日期
	OriginIncident string         `json:"origin_incident" validate:"required,max=512" label:"事由"` // 账户产生的事由
	Amount         string         `json:"amount" validate:"required" label:"账单金额"`                // 账单金额（数据库单位分，接收前端值用string避免丢失精度）
	Relation       string         `json:"relation" validate:"required,max=32" label:"关系"`         // 与对方关系,如亲戚|同事|闺蜜
	ToName         string         `json:"to_name" validate:"max=32"`                              // 对方姓名
	IsFollow       uint           `json:"is_follow" validate:"oneof=0 1" label:"关注状态"`            // 是否关注或者跟进，0不关注、1关注
	Remark         string         `json:"remark" validate:"required" label:"备注"`
}

// ToAccountBill 用户请求转为数据模型
func (abr AddAccountBillReq) ToAccountBill(uid uint64, amount uint) AccountBill {
	return AccountBill{
		UserId:         uid,
		BillDate:       time.Time(abr.BillDate),
		OriginIncident: abr.OriginIncident,
		Amount:         amount,
		Relation:       abr.Relation,
		ToName:         abr.ToName,
		IsFollow:       abr.IsFollow,
		Remark:         abr.Remark,
	}
}

// AccountBillResp 列出用户账目清单的响应数据
// 因为跟请求信息基本一致，这里直接定义别名
type AccountBillResp AddAccountBillReq
