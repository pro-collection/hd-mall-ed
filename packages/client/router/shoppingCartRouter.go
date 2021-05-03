package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/shoppingCartController"
)

func ShoppingCartRouter(router *gin.RouterGroup)  {
	shoppingCart := router.Group("/shopping_cart")
	{
		// 获取当前用户的所有购物车/临时订单
		shoppingCart.GET("/get", shoppingCartController.GetList)

		// 根据 id 获取信息
		shoppingCart.GET("/get_by_id", shoppingCartController.GetDetailByID)

		// 根据 temp order id 获取信息列表
		shoppingCart.GET("/get_by_temp_order_id", shoppingCartController.GetDetailByTempOrderId)

		// 创建购物车/临时订单
		shoppingCart.POST("/create", shoppingCartController.Create)

		// 删除购物车/临时订单
		shoppingCart.POST("/delete", shoppingCartController.Delete)

		// 更新
		shoppingCart.POST("/update", shoppingCartController.Update)
	}
}
