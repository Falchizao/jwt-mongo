package model

import (
	"github.com/Falchizao/api-golang/src/configuration/handling_err"
)

func (u *UserControllerInterface) Find(string) (*User, *handling_err.RestErr) {

	u.Encrypt_pass()
	return nil, nil
}
