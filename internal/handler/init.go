// Created on 2021/3/14.
// @author tony
// email xmgtony@gmail.com
// description 初始化handler，注入service

package handler

import (
	"apiserver-gin/internal/handler/v1"
	"apiserver-gin/internal/service"
)

var s *service.Service

func InitHandler(service *service.Service) {
	s = service
	v1.InitV1Handler(s)
}
