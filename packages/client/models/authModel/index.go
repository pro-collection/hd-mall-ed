package authModel

import "hd-mall-ed/packages/client/database"

type Auth struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (int, error) {
	var auth Auth
	err := database.DataBase.Select("id").
		Where(Auth{Username: username, Password: password}).
		First(&auth).Error
	if auth.Id > 0 {
		return auth.Id, nil
	}
	return -1, err
}

func GetAuthById(id int64) Auth {
	var auth Auth
	database.DataBase.Where("id = ?", id).First(&auth)
	return auth
}
