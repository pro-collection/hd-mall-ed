package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/authController"
)

func authRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.GET("/logout", authController.Logout)
	}
}
