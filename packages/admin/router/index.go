package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/userController"
	"hd-mall-ed/packages/common/middleware/jwtMiddleware"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// auth - 不需要授权的场景
	auth := router.Group("/api/admin")
	{
		authRouter(auth)
		// 注册用户
		auth.POST("/user/register", userController.Register)
	}

	// 需要授权的场景
	index := router.Group("/api/admin")
	index.Use(jwtMiddleware.AdminJwt())
	{
		// user
		userRouter(index)
		// category
		categoryRouter(index)
		// upload
		uploadRouter(index)
		// product
		productRouter(index)
		// productCategoryRouter
		productCategoryRouter(index)

		// 资源
		staticRouter(index)
	}

	return router
}
