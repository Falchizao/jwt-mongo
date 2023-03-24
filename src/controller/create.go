package controller

import (
	"fmt"
	"net/http"

	"github.com/Falchizao/api-golang/src/configuration/handling_err"
	"github.com/Falchizao/api-golang/src/controller/model"
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

	user := model.NewUser(user_req.Email, user_req.Password, user_req.Name, user_req.Age)

	if err := user.Create(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.String(http.StatusOK, "")
}
