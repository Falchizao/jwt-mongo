package model

import (
	"github.com/Falchizao/api-golang/src/configuration/handling_err"
)

func (u *User) Update(string) *handling_err.RestErr {

	u.Encrypt_pass()
	return nil
}
