package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Falchizao/api-golang/src/configuration/handling_err"
)

type User struct {
	Email    string
	Password string
	Name     string
	Age      int
}

func NewUser(email string, password string, name string, age int) *User {
	return &User{email, password, name, age}
}

type UserInterface interface {
	Create() *handling_err.RestErr
	Update(string) *handling_err.RestErr
	Find(string) (*User, *handling_err.RestErr)
	Delete(string) *handling_err.RestErr
}

func (u *User) Encrypt_pass() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
