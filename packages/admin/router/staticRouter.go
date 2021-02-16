package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/staticController"
)

func staticRouter(router *gin.RouterGroup) {
	static := router.Group("/static")
	//static.Use(jwtMiddleware.AdminJwt())
	{
		static.GET("/list", staticController.GetListByQuery)
		static.POST("/create", staticController.CreateStatics)
		static.POST("/update", staticController.Update)
		static.POST("/delete", staticController.Deletes)
	}
}
