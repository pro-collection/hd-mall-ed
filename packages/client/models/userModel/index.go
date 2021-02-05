package userModel

import (
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/database/tableModel"
)

type User tableModel.User

// 创建用户
func (user *User) CreateUser() (err error) {
	err = database.DataBase.Create(&user).Error
	return
}

// 查找user
func (user *User) FindUser() User {
	database.DataBase.First(user)
	return *user
}

// 通过 id 查找user
func (*User) FindUserById(id uint) (User, error) {
	resultUser := User{}
	err := database.DataBase.First(&resultUser, id).Error
	if err != nil {
		return resultUser, err
	}
	return resultUser, nil
}

// 通过用户名查找
func (*User) FindUserByName(name string) (User, error) {
	resultUser := User{}
	err := database.DataBase.Where("name = ?", name).First(&resultUser).Error
	if err != nil {
		return resultUser, err
	}
	return resultUser, nil
}

// 更新方法
func (user *User) Update(updateMap interface{}) error {
	err := database.DataBase.Model(&User{}).Where("id = ?", user.ID).Updates(updateMap).Error
	return err
}

// 更新默认地址
func (user *User) updateDefaultAddressId(addressId int) error {
	err := database.DataBase.Model(&User{}).Where("id = ?", user.ID).Update("default_address_id", addressId).Error
	return err
}
