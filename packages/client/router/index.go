package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()


	index := router.Group("/api")
	{

	}
}
