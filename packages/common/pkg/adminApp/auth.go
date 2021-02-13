package adminApp

import (
	"hd-mall-ed/packages/admin/models/adminUserModel"
	"hd-mall-ed/packages/common/pkg/utils/jwtUtil"
)

// 通过 token 获取用户信息
func (api *ApiFunction) GetUser() (adminUserModel.AdminUser, error) {
	var user adminUserModel.AdminUser

	token, err := api.C.Cookie("admin-token")
	if err != nil {
		return user, err
	}

	user, err = jwtUtil.AdminParseTokenGetUser(token)

	if err != nil {
		return user, err
	}

	return user, nil
}

// 单独只获取用户的id
func (api *ApiFunction) GetUserId() int {
	user, err := api.GetUser()
	if err != nil {
		return -1
	}
	return int(user.ID)
}
