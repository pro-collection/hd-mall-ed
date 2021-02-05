package userController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
	"hd-mall-ed/packages/client/service/userService"
	"log"
)

func UpdateUser(c *gin.Context) {
	var err error
	api := app.ApiFunction{C: c}
	user := &userModel.User{}

	// 绑定参数
	if handleBindUserHasErrorHelper(user, c) {
		return
	}

	// 校验 id 是否存在
	if handleCheckUserIdNotExistHelper(user, c) {
		return
	}

	// 校验用户名是否重复
	if handleCheckUserExistHelper(user, c) {
		return
	}

	// 更新的业务逻辑
	err = userService.UserUpdate(user)
	if err != nil {
		log.Println(err.Error())
		api.ResFail(e.Fail)
		return
	}

	api.ResponseNoData()
}
