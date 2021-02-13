package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/userController"
	"hd-mall-ed/packages/common/middleware/jwtMiddleware"
)

func userRouter(router *gin.RouterGroup) {
	user := router.Group("/user")

	// 注册用户
	user.POST("register", userController.Register)

	user.Use(jwtMiddleware.AdminJwt())
	{

	}
}
