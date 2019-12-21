package check

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HealthCheck check server start up status
func HealthCheck(c *gin.Context) {
	message := "SUCCESS"
	c.String(http.StatusOK, "\r\n"+message)
}

// LocalIPCheck check request IP is be allowed
func LocalIPCheck(c *gin.Context) {
	hostInfo := strings.Split(c.Request.Host, ":")
	ip := hostInfo[0]
	allowIps := []string{"localhost", "127.0.0.1"}
	for _, item := range allowIps {
		if ip == item {
			return
		}
	}
	c.AbortWithStatus(http.StatusForbidden)
}
