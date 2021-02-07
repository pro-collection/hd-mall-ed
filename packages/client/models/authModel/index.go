package authModel

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/models/userModel"
)


func CheckAuth(username, password string) (userModel.User, error) {
	var user userModel.User

	queryMap := map[string]interface{}{
		"name":     username,
		"password": password,
	}

	err := database.DataBase.
		Where(queryMap).
		First(&user).Error
	if user.ID > 0 {
		return user, nil
	}
	return user, err
}

func GetAuthById(id int64) userModel.User {
	var user userModel.User
	database.DataBase.Where("id = ?", id).First(&user)
	return user
}
