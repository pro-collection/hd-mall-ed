package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/productController"
)

func productRouter(router *gin.RouterGroup) {
	product := router.Group("/product")
	//product.Use(jwtMiddleware.Jwt())
	{
		// 获取列表
		product.POST("/list", productController.GetListByQuery)
	}
}
