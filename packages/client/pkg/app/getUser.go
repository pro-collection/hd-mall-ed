package app

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/client/pkg/utils"
)

// 通过 token 获取用户信息
func (api *ApiFunction) GetUser() (userModel.User, error) {
	var user userModel.User

	token, err := api.C.Cookie("token")
	if err != nil {
		return user, err
	}
	claims, err2 := utils.ParseToken(token)
	if err2 != nil {
		return user, err2
	}

	err = database.DataBase.Model(&userModel.User{}).
		Where("id = ?", claims.Id).
		First(&user).
		Error
	if err != nil {
		return user, err
	}
	return user, nil
}
