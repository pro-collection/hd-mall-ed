package userService

import "hd-mall-ed/packages/client/models/userModel"

func UserUpdate(user *userModel.User) error {
	// save 是全量更新
	updateMap := make(map[string]interface{})
	updateMap["password"] = user.Password

	return user.Update(updateMap)
}
