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
		// 创建分类
		category.POST("/create", categoryController.Create)

		// 获取所有分类列表
		category.GET("/list", categoryController.GetList)

		// 更新分类
		category.POST("/update", categoryController.Update)
	}
}
