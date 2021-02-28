package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/bannerController"
)

func bannerRouter(router *gin.RouterGroup) {
	banner := router.Group("/banner")
	{
		banner.GET("/list", bannerController.GetBannerList)
		banner.POST("/create", bannerController.Create)
		banner.POST("/update", bannerController.Update)
		banner.POST("/delete", bannerController.Delete)
	}
}
