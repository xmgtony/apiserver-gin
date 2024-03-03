package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping ping服务器状态
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "\r\nSUCCESS")
	}
}
