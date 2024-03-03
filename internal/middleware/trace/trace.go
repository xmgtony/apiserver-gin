package trace

import (
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/log"
	"apiserver-gin/tools/uuid"
	"context"
	"github.com/gin-gonic/gin"
)

// SetRequestId 用来设置和透传requestId
func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.GenUUID16()
		c.Header("X-Request-Id", requestId)

		// 设置requestId到context中，便于后面调用链的透传
		c.Set(constant.RequestId, requestId)
		c.Next()
	}
}

// RequestId 获取requestId
func RequestId() log.Valuer {
	return func(c context.Context) any {
		if rid := c.Value(constant.RequestId); rid != nil {
			return rid.(string)
		}
		return ""
	}
}
