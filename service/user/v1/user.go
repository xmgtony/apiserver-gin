package v1

import (
	"apidemo-gin/model/user"
	"apidemo-gin/pkg/errcode"
	"apidemo-gin/tools/security"
)

func Create(user *user.User) error {
	err := user.Validate()
	if err != nil {
		return errcode.New(errcode.ValidateErr, err)
	}
	if pwd, err := security.Encrypt(user.Password); err != nil {
		return err
	} else {
		user.Password = pwd
	}
	if err := user.AddUser(); err != nil {
		return err
	}
	return nil
}
