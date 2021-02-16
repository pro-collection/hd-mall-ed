package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/productCategoryController"
)

func productCategoryRouter(router *gin.RouterGroup) {
	category := router.Group("/product_category")
	//category.Use(jwtMiddleware.AdminJwt())
	{
		category.GET("/list", productCategoryController.GetList)
		category.POST("/update", productCategoryController.Update)
		category.POST("/create", productCategoryController.Create)
		category.POST("/delete", productCategoryController.Delete)
	}
}
