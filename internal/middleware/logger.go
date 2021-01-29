package middleware

import (
	"apiserver-gin/pkg/constant"
	. "apiserver-gin/pkg/log"
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

// Logger 记录每次请求的请求信息和响应信息
func Logger(c *gin.Context) {
	// 请求前
	t := time.Now()
	reqPath := c.Request.URL.Path
	reqId := c.GetString(constant.RequestId)
	method := c.Request.Method
	ip := c.ClientIP()
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		requestBody = []byte{}
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	Log.Info("New request start",
		zap.String(constant.RequestId, reqId),
		zap.String("host", ip),
		zap.String("path", reqPath),
		zap.String("method", method),
		zap.String("body", string(requestBody)))

	c.Next()
	// 请求后
	latency := time.Since(t)
	Log.Info("New request end",
		zap.String(constant.RequestId, reqId),
		zap.String("host", ip),
		zap.String("path", reqPath),
		zap.Duration("cost", latency))
}
