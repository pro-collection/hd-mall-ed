package userController

import (
	"hd-mall-ed/packages/admin/models/adminUserModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

// @title           判定用户是否存在
// @description
// @auth            晴小篆  331393627@qq.com
// @param           user *adminUserModel.AdminUser, api *adminApp.ApiFunction
// @return          bool 是否存在， true 用户名存在；false 用户名不存在
func handleCheckUserNameIsExistHelper(user *adminUserModel.AdminUser, api *adminApp.ApiFunction) bool {
	resultUser, _ := user.FindUserByName(user.Name)
	if resultUser.ID > 0 {
		api.ResFail(e.UserNameExist)
		return true
	}
	return false
}
