package middleware

import (
	"apiserver-gin/pkg/log"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// ReqLogger 记录每次请求的请求信息和响应信息
func ReqLogger(c *gin.Context) {
	// 请求前
	t := time.Now()
	reqPath := c.Request.URL.Path
	method := c.Request.Method
	ip := c.ClientIP()
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		requestBody = []byte{}
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	log.WithCtx(c).Info(fmt.Sprintf("host:%s %s %s start", ip, method, reqPath), "body", string(requestBody))

	c.Next()
	// 请求后
	latency := time.Since(t)
	log.WithCtx(c).Info(fmt.Sprintf("host:%s %s %s end", ip, method, reqPath), "cost", latency)
}
