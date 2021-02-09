package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/authController"
)

func authRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		// 登录接口
		auth.POST("/login", authController.GetAuth)

		// 退出登录
		auth.GET("/logout", authController.Logout)
	}
}
