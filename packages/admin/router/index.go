package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/api/admin")
	{
		// auth
		authRouter(index)
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
	}

	return router
}
