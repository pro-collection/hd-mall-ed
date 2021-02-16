package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/userController"
	"hd-mall-ed/packages/common/middleware/jwtMiddleware"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// 不鉴权
	auth := router.Group("/api")
	{
		// 鉴权登录的相关接口
		authRouter(auth)

		auth.POST("/user/create", userController.CreateUser)
	}

	// 需要鉴权的接口
	index := router.Group("/api")
	index.Use(jwtMiddleware.Jwt())
	{
		// user 相关的接口
		userRouter(index)

		// address 相关接口
		addressRouter(index)

		// category 商品分类
		categoryRouter(index)

		// product
		productRouter(index)

		productCategoryRouter(index)

		staticRouter(index)
	}

	return router
}
