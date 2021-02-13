package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
)

// 绑定user信息， 以及参数校验
func handleBindUserHasErrorHelper(user *userModel.User, c *gin.Context) bool {
	err := c.ShouldBindJSON(user)
	api := app.ApiFunction{C: c}
	if err != nil {
		api.ResFail(e.InvalidParams)
		return true
	}
	return false
}

// 判断该用户是否 【不存在】， 更新用户的时候判定
// 不存在的时候， 返回true
// 存在用户的时候， 返回 false
func handleCheckUserIdNotExistHelper(user *userModel.User, c *gin.Context) bool {
	api := app.ApiFunction{C: c}

	if user.ID <= 0 {
		api.ResFail(e.NotFoundId)
		return true
	}

	return false
}

// 判断用户是否存在
func handleCheckUserExistHelper(user *userModel.User, c *gin.Context) bool {
	api := app.ApiFunction{C: c}
	resultUser, _ := user.FindUserByName(user.Name)
	if resultUser.ID > 0 {
		api.ResFail(e.UserNameExist)
		return true
	}
	return false
}

func handleUserUpdateHelper(user *userModel.User) error {
	// save 是全量更新
	updateMap := make(map[string]interface{})

	// 更新密码
	if !funk.IsEmpty(user.Password) {
		updateMap["password"] = user.Password
	}

	if funk.IsEmpty(updateMap) {
		return nil
	}

	return user.Update(updateMap)
}

