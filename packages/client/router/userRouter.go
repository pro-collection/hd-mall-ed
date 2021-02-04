package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/userController"
)

func userRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/create", userController.CreateUser)
		user.POST("/update", userController.UpdateUser)
		user.GET("/info", userController.GetUserInfo)
	}
}
