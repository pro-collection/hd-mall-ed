package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/productController"
	clientProductController "hd-mall-ed/packages/client/controller/productController"
)

func productRouter(router *gin.RouterGroup) {
	product := router.Group("/product")
	//product.Use(jwtMiddleware.Jwt())
	{
		// 获取列表
		product.POST("/list", productController.GetListByQuery)

		// 获取默认的 category primary list
		product.GET("/primary", clientProductController.GetPrimaryCategoryProductList)

		// 获取discount
		product.GET("/discount", clientProductController.GetDiscountProduct)
	}
}
