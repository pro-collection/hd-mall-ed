package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/userController"
	"hd-mall-ed/packages/client/middleware/jwtMiddleware"
)

func userRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.Use(jwtMiddleware.Jwt())
	{
		user.POST("/create", userController.CreateUser)
		user.POST("/update", userController.UpdateUser)
		user.GET("/info", userController.GetUserInfo)
	}
}
