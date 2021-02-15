package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/productController"
	"hd-mall-ed/packages/common/middleware/jwtMiddleware"
)


func productRouter(router *gin.RouterGroup) {
	product := router.Group("/product")
	product.Use(jwtMiddleware.AdminJwt())
	{
		// 获取列表
		product.POST("/list", productController.GetListByQuery)

		// 创建一个商品
		product.POST("/create")
	}
}
