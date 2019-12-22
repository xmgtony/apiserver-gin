package middleware

import (
	"apidemo-gin/pkg/constant"
	. "apidemo-gin/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Logger 记录每次请求的请求信息和响应信息
func Logger(c *gin.Context) {
	// 请求前
	t := time.Now()
	reqPath := c.Request.URL.Path
	Log.Info("request start", zap.String(constant.RequestId, c.GetString(constant.RequestId)), zap.String("path", reqPath))
	c.Next()
	// 请求后
	latency := time.Since(t)
	Log.Info("request end", zap.String(constant.RequestId, c.GetString(constant.RequestId)), zap.String("path", reqPath), zap.Duration("cost", latency))
}
