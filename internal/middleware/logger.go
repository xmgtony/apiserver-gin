package middleware

import (
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/log"
	"bytes"
	"github.com/gin-gonic/gin"
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

	log.Info("New request start",
		log.Pair(constant.RequestId, reqId),
		log.Pair("host", ip),
		log.Pair("host", ip),
		log.Pair("path", reqPath),
		log.Pair("method", method),
		log.Pair("body", string(requestBody)))

	c.Next()
	// 请求后
	latency := time.Since(t)
	log.Info("New request end",
		log.Pair(constant.RequestId, reqId),
		log.Pair("host", ip),
		log.Pair("path", reqPath),
		log.Pair("cost", latency))
}
