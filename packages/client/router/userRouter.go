package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/userController"
)

func userRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	//user.Use(jwtMiddleware.Jwt())
	{
		user.POST("/update", userController.UpdateUser)
		user.GET("/info", userController.GetUserInfo)
	}
}
