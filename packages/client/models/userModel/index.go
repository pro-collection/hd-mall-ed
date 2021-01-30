package userModel

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/database/tableModal"
)

type User tableModal.User

// 创建用户
func (user *User) CreateUser() (err error) {
	err = database.DataBase.Create(&user).Error
	return
}

// 查找user
func (user *User) FindUser() User {
	database.DataBase.First(&user)
	return *user
}

// 通过 id 查找user
func (*User) findUserById(id uint) (User, error) {
	resultUser := User{}
	err := database.DataBase.First(&resultUser, id).Error
	if err != nil {
		return resultUser, err
	}
	return resultUser, nil
}


