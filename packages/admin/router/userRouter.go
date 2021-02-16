package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/userController"
)

func userRouter(router *gin.RouterGroup) {
	user := router.Group("/user")

	// 注册用户
	//user.POST("register", userController.Register)

	//user.Use(jwtMiddleware.AdminJwt())
	{
		// 获取用户信息
		user.GET("/info", userController.GetUserInfo)

		// 更新用户信息
		user.POST("/update", userController.Update)
	}
}
