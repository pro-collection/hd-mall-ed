package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/userController"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/api")
	{
		userRouter := index.Group("/user")
		{
			userRouter.POST("/create", userController.CreateUser)
			userRouter.POST("/update", userController.UpdateUser)
		}
	}

	return router
}
