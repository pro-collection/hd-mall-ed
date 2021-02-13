package userController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func GetUserInfo(c *gin.Context) {
	api := adminApp.ApiInit(c)
	user, err := api.GetUser()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}
	api.Response(user)
}
