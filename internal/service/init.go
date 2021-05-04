// Created on 2021/3/12.
// @author tony
// email xmgtony@gmail.com
// description service

package service

type Service struct {
	Us UserService
}

func InitService() *Service {
	return &Service{Us: UserService{}}
}
