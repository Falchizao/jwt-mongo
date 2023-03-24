package controller

import (
	"github.com/Falchizao/api-golang/src/configuration/handling_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := handling_err.NewBadRequestError("Erro endpoint")
	c.JSON(err.Code, err)
}
