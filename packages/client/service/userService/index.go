package userService

import (
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/client/models/userModel"
)

func UserUpdate(user *userModel.User) error {
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
