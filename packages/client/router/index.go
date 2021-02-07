package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/userController"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/api")
	{
		// auth
		index.POST("/auth", userController.GetAuth)

		// user 相关的接口
		userRouter(index)
	}

	return router
}
