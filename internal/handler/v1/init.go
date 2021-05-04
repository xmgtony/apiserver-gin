// Created on 2021/3/14.
// @author tony
// email xmgtony@gmail.com
// description 初始化handler，注入service

package v1

import "apiserver-gin/internal/service"

var s *service.Service

func InitV1Handler(service *service.Service) {
	s = service
}
