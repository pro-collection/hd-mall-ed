package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/productCategoryController"
)

func productCategoryRouter(router *gin.RouterGroup) {
	category := router.Group("/product_category")
	{
		category.GET("/list", productCategoryController.GetList)
	}
}
