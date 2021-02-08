package authModel

import (
	"encoding/json"
	"hd-mall-ed/packages/client/config/cache"
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/models/userModel"
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
	// todo 这个地方接受参数的时候有点儿问题
	userJsonString, err := cache.Manager.Get(strconv.Itoa(int(id)))
	if err != nil {
		return user
	}

	_ = json.Unmarshal([]byte(userJsonString.(string)), &user)

	return user
}
