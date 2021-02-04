package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/api")
	{
		// user 相关的接口
		userRouter(index)
	}

	return router
}
