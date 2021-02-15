package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/api/admin")
	{
		authRouter(index)
		userRouter(index)
		categoryRouter(index)
		uploadRouter(index)
	}

	return router
}
