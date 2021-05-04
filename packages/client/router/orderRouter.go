package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/orderController"
)

func orderRouter(router *gin.RouterGroup)  {
	order := router.Group("/order")
	{
		order.POST("/create", orderController.Create)
	}
}
