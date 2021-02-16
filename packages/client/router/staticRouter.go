package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/staticController"
)

func staticRouter(router *gin.RouterGroup) {
	static := router.Group("/static")
	{
		static.GET("/list", staticController.GetListByQuery)
	}
}
