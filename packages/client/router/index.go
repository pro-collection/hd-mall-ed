package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/api")
	{
		// 鉴权登录的相关接口
		authRouter(index)

		// user 相关的接口
		userRouter(index)

		// address 相关接口
		addressRouter(index)

		// category 商品分类

	}

	return router
}
