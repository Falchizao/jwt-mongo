package routes

import (
	"github.com/Falchizao/api-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {

	r.GET("/user/:userId", controller.FindUser)
	r.GET("/email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)

}
