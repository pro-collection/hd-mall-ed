package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/bannerController"
)

func bannerRouter(router *gin.RouterGroup) {
	banner := router.Group("/banner")
	{
		banner.GET("/list", bannerController.GetBannerList)
	}
}
