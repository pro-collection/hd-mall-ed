package userController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
)

// 创建用户信息
func CreateUser(c *gin.Context) {
	var err error
	api := app.ApiFunction{C: c}
	user := &userModel.User{}

	// 绑定user信息， 以及参数校验
	if handleBindUserHasErrorHelper(user, c) {
		return
	}

	if user.Name == "" {
		api.ResFail(e.InvalidParams)
		return
	}

	// 判断用户是否存在
	if handleCheckUserExistHelper(user, c) {
		return
	}

	err = user.CreateUser()
	if err != nil {
		api.ResFail(e.CreateUserFail)
		return
	}

	api.ResponseNoData()
}
