package router

import "github.com/gin-gonic/gin"

func categoryRouter(router *gin.RouterGroup) {
	category := router.Group("/category")
	{
		category.GET("/list")
	}
}
