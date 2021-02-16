package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/uploadController"
)

func uploadRouter(router *gin.RouterGroup) {
	upload := router.Group("/upload")
	//upload.Use(jwtMiddleware.AdminJwt())
	{
		upload.POST("", uploadController.Upload)
	}
}
