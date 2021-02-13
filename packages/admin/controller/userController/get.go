package userController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/common/pkg/adminApp"
)

func GetUserInfo(c *gin.Context) {
	api := adminApp.ApiInit(c)
	user, _ := api.GetUser()
	api.Response(user)
}
