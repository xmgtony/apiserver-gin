// Package reply 目前只处理json格式的响应，如果是渲染模板等，请提供其他方法，比如Render
package reply

import (
	"apiserver-gin/internal/base/constant"
	"apiserver-gin/internal/base/errcode"
	"apiserver-gin/pkg/xerrors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiResponse 代表一个响应给客户端的消息结构，包括错误码，错误消息，响应数据
type ApiResponse struct {
	Rid     string      `json:"rid"`      // 请求的唯一ID，requestId的缩写
	ErrCode int         `json:"err_code"` // 错误码，0表示无错误
	Msg     string      `json:"msg"`      // 提示信息
	Data    interface{} `json:"data"`     // 响应数据，一般从这里前端从这个里面取出数据展示
}

// PageResult 分页结果
type PageResult struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}

// JSON 发送json格式的数据
func JSON(c *gin.Context, err error, data any) {
	if err != nil {
		slog.ErrorContext(c, "error response", "error", err)
	}
	errCode, message := xerrors.DecodeErr(err)
	// 如果code != 0, 失败的话 返回http状态码400（一般也可以全部返回200）
	// 返回400 更严谨一些，个人接触的项目中大部分都是400。
	var httpStatus int
	if errCode != errcode.Success {
		httpStatus = http.StatusBadRequest
	} else {
		httpStatus = http.StatusOK
	}
	c.JSON(httpStatus, ApiResponse{
		Rid:     c.GetString(constant.TraceID),
		ErrCode: errCode,
		Msg:     message,
		Data:    data,
	})
}

// Success 响应成功的JSON数据
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResponse{
		Rid:     c.GetString(constant.TraceID),
		ErrCode: errcode.Success,
		Msg:     "success",
		Data:    data,
	})
}

func Fail(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Rid:     c.GetString(constant.TraceID),
			ErrCode: errcode.Unknown,
			Msg:     "unknown error",
			Data:    nil,
		})
	}
	if err != nil {
		slog.ErrorContext(c, "error response", "error", err)
		errCode, message := xerrors.DecodeErr(err)
		c.JSON(http.StatusBadRequest, ApiResponse{
			Rid:     c.GetString(constant.TraceID),
			ErrCode: errCode,
			Msg:     message,
			Data:    nil,
		})
	}
}

func Page[T any](c *gin.Context, totalCount int64, list []T) {
	page := &PageResult{
		Total: totalCount,
		List:  list,
	}
	Success(c, page)
}
