package authModel

import (
	"encoding/json"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/common/config/cache"
	"hd-mall-ed/packages/common/database"
	"strconv"
)

type Auth struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (int, error) {
	var user userModel.User
	queryMap := map[string]interface{}{
		"name":     username,
		"password": password,
	}

	err := database.DataBase.Model(&userModel.User{}).
		Where(queryMap).
		First(&user).Error

	// 缓存
	userBytes, _ := json.Marshal(user)
	err = cache.Manager.Set(strconv.Itoa(int(user.ID)), string(userBytes), nil)
	if err != nil {
		return -1, err
	}

	if user.ID > 0 {
		return int(user.ID), nil
	}
	return -1, err
}

func GetAuthById(id int64) userModel.User {
	var user userModel.User
	// 获取鉴权
	userJsonString, err := cache.Manager.Get(strconv.Itoa(int(id)))

	// 没有获取到缓存， 或者缓存失效了
	if userJsonString == "" {
		if err := database.DataBase.Model(&userModel.User{}).Where("id = ?", id).First(&user).Error; err != nil {
			userBytes, _ := json.Marshal(user)
			_ = cache.Manager.Set(strconv.Itoa(int(user.ID)), string(userBytes), nil)
			return user
		}
	}

	if err != nil {
		return user
	}

	_ = json.Unmarshal([]byte(userJsonString.(string)), &user)

	return user
}
