package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/productController"
)

func productRouter(router *gin.RouterGroup) {
	product := router.Group("/product")
	//product.Use(jwtMiddleware.AdminJwt())
	{
		// 获取列表
		product.POST("/list", productController.GetListByQuery)

		// 创建一个商品
		product.POST("/create", productController.Create)

		// 更新一个商品
		product.POST("/update", productController.Update)

		// 删除
		product.POST("/delete", productController.Delete)

		// 通过 id 获取一个 详情数据
		product.GET("/get_detail", productController.GetDetail)
	}
}
