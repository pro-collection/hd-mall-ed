package authModel

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/models/userModel"
)

type Auth struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (int, error) {
	var auth Auth
	queryMap := map[string]interface{}{
		"name":     username,
		"password": password,
	}

	err := database.DataBase.Model(&userModel.User{}).
		Select("id").
		Where(queryMap).
		First(&auth).Error

	if auth.Id > 0 {
		return auth.Id, nil
	}
	return -1, err
}

func GetAuthById(id int64) userModel.User {
	var user userModel.User
	// todo 这个地方接受参数的时候有点儿问题
	database.DataBase.Where("id = ?", id).First(&user)
	return user
}
