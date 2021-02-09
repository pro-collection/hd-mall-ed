package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/addressController"
	"hd-mall-ed/packages/client/middleware/jwtMiddleware"
)

func addressRouter(router *gin.RouterGroup)  {
	address := router.Group("/address")
	address.Use(jwtMiddleware.Jwt())
	{
		// 获取所有列表
		address.GET("/list", addressController.GetAddressList)
		// 创建地址
		address.POST("/create", addressController.Create)
		// 通过 id 查询具体的地址
		address.GET("/get", addressController.GetAddressById)
		// 修改地址
		// 删除地址
	}
}
