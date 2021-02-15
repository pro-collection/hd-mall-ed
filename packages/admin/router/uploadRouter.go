package router

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/controller/uploadController"
	"hd-mall-ed/packages/common/middleware/jwtMiddleware"
)

func uploadRouter(router *gin.RouterGroup) {
	upload := router.Group("/upload")

	upload.Use(jwtMiddleware.AdminJwt())
	{
		upload.POST("", uploadController.Upload)
	}
}
