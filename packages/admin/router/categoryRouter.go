package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/categoryController"
	"hd-mall-ed/packages/common/middleware/jwtMiddleware"
)

func categoryRouter(router *gin.RouterGroup) {
	category := router.Group("/category")
	category.Use(jwtMiddleware.AdminJwt())
	{
		category.POST("/create", categoryController.Create)

		// 获取所有列表
		category.GET("/list", categoryController.GetList)
	}
}
