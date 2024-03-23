package middleware

import (
	"apiserver-gin/pkg/log"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

// ReqLogger 记录每次请求的请求信息和响应信息
func ReqLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		t := time.Now()
		reqPath := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()
		rawQuery := c.Request.URL.RawQuery
		// JSON和FORM表单打印请求的Body, 其他内容类型，比如文件上传不打印
		var requestBody string
		contentType := c.GetHeader("Content-Type")
		if contentType != "" &&
			(strings.HasPrefix(contentType, "application/json") ||
				strings.HasPrefix(contentType, "application/x-www-form-urlencoded")) {
			requestBody, err := io.ReadAll(c.Request.Body)
			if err != nil {
				requestBody = []byte{}
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		log.WithCtx(c).Info(fmt.Sprintf("host:%s %s %s start", ip, method, reqPath), "query", rawQuery, "body", requestBody)

		c.Next()
		// 请求后
		latency := time.Since(t).Microseconds()
		log.WithCtx(c).Info(fmt.Sprintf("host:%s %s %s end", ip, method, reqPath), "cost/us", latency)
	}
}
