package adminAuthModel

import (
	"encoding/json"
	"hd-mall-ed/packages/common/config/cache"
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/models/adminUserModel"
)

func CheckAuth(username, password string) (int, error) {
	var user adminUserModel.AdminUser
	queryMap := map[string]interface{}{
		"name":     username,
		"password": password,
	}

	err := database.DataBase.Model(&adminUserModel.AdminUser{}).
		Where(queryMap).
		First(&user).Error

	// 缓存
	userBytes, _ := json.Marshal(user)
	err = cache.Manager.Set(cache.AminCacheKey(int(user.ID)), string(userBytes), nil)
	if err != nil {
		return -1, err
	}

	if user.ID > 0 {
		return int(user.ID), nil
	}
	return -1, err
}

func GetAuthById(id int64) adminUserModel.AdminUser {
	var user adminUserModel.AdminUser
	// 获取鉴权
	userJsonString, err := cache.Manager.Get(cache.AminCacheKey(int(id)))
	if err != nil {
		return user
	}

	_ = json.Unmarshal([]byte(userJsonString.(string)), &user)

	return user
}
