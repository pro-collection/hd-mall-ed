package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/controller/addressController"
)

func addressRouter(router *gin.RouterGroup)  {
	address := router.Group("/address")
	{
		// 获取所有列表
		address.GET("/list", addressController.GetAddressList)
	}
}
