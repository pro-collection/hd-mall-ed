package app

import (
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/common/pkg/utils/jwtUtil"
)

// 通过 token 获取用户信息
func (api *ApiFunction) GetUser() (userModel.User, error) {
	var user userModel.User

	token, err := api.C.Cookie("token")
	if err != nil {
		return user, err
	}

	user, err = jwtUtil.ParseTokenGetUser(token)

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

func (api *ApiFunction) SetCookie(key, value string, maxAge int) {
	api.C.SetCookie(key, value, maxAge, "", "", false, false)
}

func (api *ApiFunction) RemoveCookie(key string) {
	api.C.SetCookie(key, "", 0, "", "", false, false)
}
