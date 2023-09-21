package middleware

import (
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/log"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		requestBody = []byte{}
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	log.Info(fmt.Sprintf("%s %s start", method, reqPath),
		log.Pair(constant.RequestId, reqId),
		log.Pair("host", ip),
		log.Pair("host", ip),
		log.Pair("path", reqPath),
		log.Pair("method", method),
		log.Pair("body", string(requestBody)))

	c.Next()
	// 请求后
	latency := time.Since(t)
	log.Info(fmt.Sprintf("%s %s end", method, reqPath),
		log.Pair(constant.RequestId, reqId),
		log.Pair("host", ip),
		log.Pair("path", reqPath),
		log.Pair("cost", latency))
}
