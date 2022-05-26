// Created on 2022/5/27.
// @author tony
// email xmgtony@gmail.com
// description 账目清单handler

package accountbill

import (
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/ecode"
	"apiserver-gin/pkg/response"
	jtime "apiserver-gin/pkg/time"
	"github.com/gin-gonic/gin"
	"time"
)

// Handler 账目清单handler，从分层来讲，这里已经是最外层，只要提供实例在router中使用
// 所以这里没有定义接口，而是直接使用struct来组织多个handler func
type Handler struct {
	accountBillServ service.AccountBillService
}

func NewAccountBillHandler(_accountBillServ service.AccountBillService) *Handler {
	return &Handler{
		accountBillServ: _accountBillServ,
	}
}

func (abh *Handler) AddAccountBill() gin.HandlerFunc {
	return func(c *gin.Context) {
		addAccountBillReq := model.AddAccountBillReq{}
		if err := c.ShouldBind(&addAccountBillReq); err != nil {
			response.JSON(c, errors.WithCode(ecode.ValidateErr, err.Error()), nil)
			return
		}
		uid := c.GetInt64(constant.UserID)
		// 组织model信息
		accountBill := model.AccountBill{
			UserId:         uint64(uid),
			BillDate:       time.Time(addAccountBillReq.BillDate),
			OriginIncident: addAccountBillReq.OriginIncident,
			Amount:         0,
			Relation:       "",
			ToName:         "",
			IsFollow:       0,
			Remark:         "",
		}
		err := abh.accountBillServ.Save(c, &accountBill)
		if err != nil {
			response.JSON(c, errors.Wrap(err, ecode.RecordCreateErr, "保存账目清单信息失败"), nil)
			return
		}
		response.JSON(c, nil, accountBill)
	}
}

func (abh *Handler) GetAccountBillList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用auth中间件的接口，一定能解析出来，否则中间件会响应错误
		uid := c.GetInt64(constant.UserID)
		bills, err := abh.accountBillServ.SelectListByUserId(c, uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, ecode.NotFoundErr, "查询错误，未找到记录"), nil)
			return
		}
		respBills := make([]model.AccountBillResp, 0)
		for _, bill := range bills {
			respBill := model.AccountBillResp{
				BillDate:       jtime.JsonTime(bill.BillDate),
				OriginIncident: bill.OriginIncident,
				Amount:         "0.00",
				Relation:       "",
				ToName:         "",
				IsFollow:       0,
				Remark:         "",
			}
			respBills = append(respBills, respBill)
		}
		response.JSON(c, nil, respBills)
		return
	}
}
