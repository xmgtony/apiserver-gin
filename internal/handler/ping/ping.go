package ping

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

// Ping ping服务器状态
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "\r\nSUCCESS")
	}
}

// isLocalIP 检查请求的ip是否是本地ip
func isLocalIP(host string) bool {
	ip, _, err := net.SplitHostPort(host)
	if err != nil {
		return false
	}
	allowIps := []string{"localhost", "127.0.0.1"}
	for _, item := range allowIps {
		if ip == item {
			return true
		}
	}
	return false
}
