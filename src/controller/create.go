package controller

import (
	"fmt"

	"github.com/Falchizao/api-golang/src/configuration/handling_err"
	"github.com/Falchizao/api-golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user_req request.UserRequest

	if err := c.ShouldBindJSON(&user_req); err != nil {
		restErr := handling_err.NewBadRequestError(fmt.Sprintf("Campos incorretos, erro=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(user_req)

}
