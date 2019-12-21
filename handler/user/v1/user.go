package v1

import (
	usermodel "apidemo-gin/model/user"
	"apidemo-gin/pkg/response"
	"github.com/gin-gonic/gin"
)

// Get get user info by user name
func Get(c *gin.Context) {
	username := c.Param("name")
	user, err := usermodel.GetUser(username)
	if err != nil {
		response.SendJson(c, err, nil)
		return
	}
	response.SendJson(c, nil, user)
}
