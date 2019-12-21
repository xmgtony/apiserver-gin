package v1

import (
	usermodel "apidemo-gin/model/user"
	"apidemo-gin/pkg/errcode"
	"apidemo-gin/pkg/response"
	"apidemo-gin/pkg/time"
	userservicev1 "apidemo-gin/service/user/v1"
	"github.com/gin-gonic/gin"
)

// CreateReq receive create user request
type CreateReq struct {
	Name     string        `json:"name"`
	Password string        `json:"password"`
	Birthday time.JsonTime `json:"birthday"`
}

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

// Create add user
func Create(c *gin.Context) {
	createReq := CreateReq{}
	err := c.ShouldBindJSON(&createReq)
	if err != nil {
		response.SendJson(c, errcode.BindingErr, nil)
		return
	}

	user := usermodel.User{
		Name:     createReq.Name,
		Password: createReq.Password,
	}

	err = userservicev1.Create(&user)
	if err != nil {
		response.SendJson(c, err, nil)
		return
	}
	response.SendJson(c, nil, nil)
	return
}
