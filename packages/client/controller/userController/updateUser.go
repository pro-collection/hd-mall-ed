package userController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/userModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"log"
)

func UpdateUser(c *gin.Context) {
	var err error
	var user userModel.User
	api := app.ApiFunction{C: c}
	user = userModel.User{}

	// 绑定参数
	if handleBindUserHasErrorHelper(&user, c) {
		return
	}

	user.ID = uint(api.GetUserId())

	// 校验 id 是否存在
	if handleCheckUserIdNotExistHelper(&user, c) {
		return
	}

	// 校验用户名是否重复
	if handleCheckUserExistHelper(&user, c) {
		return
	}

	// 更新的业务逻辑
	err = handleUserUpdateHelper(&user)
	if err != nil {
		log.Println(err.Error())
		api.ResFail(e.Fail)
		return
	}

	api.ResponseNoData()
}
